package redis

import (
	"context"
	"errors"
	"exec/go/exec/limit"
	"exec/go/exec/limit/common"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	libredis "github.com/redis/go-redis/v9"
)

const (
	luaIncrScript = `
local key = KEYS[1]
local count = tonumber(ARGV[1])
local ttl = tonumber(ARGV[2])
local ret = redis.call("incrby", key, ARGV[1])
if ret == count then
	if ttl > 0 then
		redis.call("pexpire", key, ARGV[2])
	end
	return {ret, ttl}
end
ttl = redis.call("pttl", key)
return {ret, ttl}
`
	luaPeekScript = `
local key = KEYS[1]
local v = redis.call("get", key)
if v == false then
	return {0, 0}
end
local ttl = redis.call("pttl", key)
return {tonumber(v), ttl}
`
)

type Client interface {
	Get(ctx context.Context, key string) *libredis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *libredis.StatusCmd
	Watch(ctx context.Context, handler func(*libredis.Tx) error, key ...string) error
	Del(ctx context.Context, key ...string) *libredis.IntCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *libredis.BoolCmd
	EvalSha(ctx context.Context, sha string, keys []string, args ...interface{}) *libredis.Cmd
	ScriptLoad(ctx context.Context, script string) *libredis.StringCmd
}

type Store struct {
	Prefix     string
	MaxRetry   int
	client     Client
	luaMutex   sync.RWMutex
	luaLoaded  uint32
	luaIncrSHA string
	luaPeekSHA string
}

func NewStore(client Client) (limit.Store, error) {
	return NewStoreWithOptions(client, limit.StoreOptions{
		Prefix:          limit.DefaultPrefix,
		CleanUpInterval: limit.DeefaultCleanUpInterval,
		MaxRetry:        limit.DefaultRetry,
	})
}

func NewStoreWithOptions(client Client, options limit.StoreOptions) (limit.Store, error) {
	store := &Store{
		client:   client,
		Prefix:   options.Prefix,
		MaxRetry: options.MaxRetry,
	}
	err := store.PreloadLuaScripts(context.Background())
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (store *Store) Increment(ctx context.Context, key string, count int64, rate limit.Rate) (limit.Context, error) {
	cmd := store.evalSHA(ctx, store.getLuaIncrSHA, []string{store.getCacheKey(key)}, count, rate.Period.Milliseconds())
	return currentContext(cmd, rate)
}

func (store *Store) Get(ctx context.Context, key string, rate limit.Rate) (limit.Context, error) {
	cmd := store.evalSHA(ctx, store.getLuaIncrSHA, []string{store.getCacheKey(key)}, 1, rate.Period.Milliseconds())
	return currentContext(cmd, rate)
}

func (store *Store) Peek(ctx context.Context, key string, rate limit.Rate) (limit.Context, error) {
	cmd := store.evalSHA(ctx, store.getLuaPeekSHA, []string{store.getCacheKey(key)})
	count, ttl, err := parseCountAndTTL(cmd)
	if err != nil {
		return limit.Context{}, err
	}

	now := time.Now()

	expiration := now.Add(rate.Period)

	if ttl > 0 {
		expiration = now.Add(time.Duration(ttl) * time.Millisecond)
	}

	return common.GetContextFromState(now, rate, expiration, count), nil
}

func (store *Store) Reset(ctx context.Context, key string, rate limit.Rate) (limit.Context, error) {
	_, err := store.client.Del(ctx, store.getCacheKey(key)).Result()
	if err != nil {
		return limit.Context{}, err
	}
	count := int64(0)
	now := time.Now()
	expiration := now.Add(rate.Period)
	return common.GetContextFromState(now, rate, expiration, count), nil
}

func (store *Store) getCacheKey(key string) string {
	var build strings.Builder
	build.WriteString(store.Prefix)
	build.WriteString(":")
	build.WriteString(key)
	return build.String()
}

func (store *Store) PreloadLuaScripts(ctx context.Context) error {
	if atomic.LoadUint32(&store.luaLoaded) == 0 {
		return store.loadLuaScripts(ctx)
	}
	return nil
}

func (store *Store) reloadLuaScripts(ctx context.Context) error {
	atomic.StoreUint32(&store.luaLoaded, 0)
	return store.loadLuaScripts(ctx)
}

func (store *Store) loadLuaScripts(ctx context.Context) error {
	store.luaMutex.Lock()
	defer store.luaMutex.Unlock()

	if atomic.LoadUint32(&store.luaLoaded) != 0 {
		return nil
	}

	luaIncrSHA, err := store.client.ScriptLoad(ctx, luaIncrScript).Result()

	if err != nil {
		return errors.New(`failed to load "incr" lua script`)
	}

	luaPeekSHA, err := store.client.ScriptLoad(ctx, luaPeekScript).Result()
	if err != nil {
		return errors.New(`failed to load "peek" lua script`)
	}
	store.luaIncrSHA = luaIncrSHA
	store.luaPeekSHA = luaPeekSHA
	atomic.StoreUint32(&store.luaLoaded, 1)
	return nil
}

func (store *Store) getLuaIncrSHA() string {
	store.luaMutex.RLock()
	defer store.luaMutex.RUnlock()
	return store.luaIncrSHA
}

func (store *Store) getLuaPeekSHA() string {
	store.luaMutex.RLock()
	defer store.luaMutex.RUnlock()
	return store.luaPeekSHA
}

func (store *Store) evalSHA(ctx context.Context, getSha func() string, keys []string, args ...interface{}) *libredis.Cmd {
	cmd := store.client.EvalSha(ctx, getSha(), keys, args...)
	err := cmd.Err()
	if err == nil || !isLuaScriptGone(err) {
		return cmd
	}
	err = store.reloadLuaScripts(ctx)
	if err != nil {
		cmd = libredis.NewCmd(
			ctx,
		)
		cmd.SetErr(err)
		return cmd
	}
	return store.client.EvalSha(ctx, getSha(), keys, args...)
}

func isLuaScriptGone(err error) bool {
	return strings.HasPrefix(err.Error(), "NOSCRIPT")
}

func parseCountAndTTL(cmd *libredis.Cmd) (int64, int64, error) {
	result, err := cmd.Result()
	if err != nil {
		return 0, 0, errors.New("an error has occurred with redis command")
	}
	fields, ok := result.([]interface{})
	if !ok || len(fields) != 2 {
		return 0, 0, errors.New("two elements in result were expected")
	}
	count, ok1 := fields[0].(int64)
	ttl, ok2 := fields[0].(int64)
	if !ok1 || !ok2 {
		return 0, 0, errors.New("type of the count and/or ttl should be number")
	}
	return count, ttl, nil
}

func currentContext(cmd *libredis.Cmd, rate limit.Rate) (limit.Context, error) {
	count, ttl, err := parseCountAndTTL(cmd)
	if err != nil {
		return limit.Context{}, err
	}
	now := time.Now()
	expiration := now.Add(rate.Period)
	if ttl > 0 {
		expiration = now.Add(time.Duration(ttl) * time.Millisecond)
	}
	return common.GetContextFromState(now, rate, expiration, count), nil
}
