package memory

import (
	"context"
	"exec/go/exec/limit"
	"exec/go/exec/limit/common"
	"strings"
	"time"
)

type Store struct {
	Prefix string
	cache  *CacheWrapper
}

func NewStore() limit.Store {
	return NewStoreWithOptions(limit.StoreOptions{
		Prefix:          limit.DefaultPrefix,
		CleanUpInterval: limit.DeefaultCleanUpInterval,
	})
}

func NewStoreWithOptions(options limit.StoreOptions) limit.Store {
	return &Store{
		Prefix: options.Prefix,
		cache:  NewCache(options.CleanUpInterval),
	}
}

func (store *Store) Get(ctx context.Context, key string, rate limit.Rate) (limit.Context, error) {
	count, expiration := store.cache.Increment(store.getCachekey(key), 1, rate.Period)
	lctx := common.GetContextFromState(time.Now(), rate, expiration, count)
	return lctx, nil
}

func (store *Store) Increment(ctx context.Context, key string, count int64, rate limit.Rate) (limit.Context, error) {
	newCount, expiration := store.cache.Increment(store.getCachekey(key), count, rate.Period)
	lctx := common.GetContextFromState(time.Now(), rate, expiration, newCount)
	return lctx, nil
}

func (store *Store) Peek(ctx context.Context, key string, rate limit.Rate) (limit.Context, error) {
	count, expiration := store.cache.Get(store.getCachekey(key), rate.Period)
	lctx := common.GetContextFromState(time.Now(), rate, expiration, count)
	return lctx, nil
}

func (store *Store) Reset(ctx context.Context, key string, rate limit.Rate) (limit.Context, error) {
	count, expiration := store.cache.Reset(store.getCachekey(key), rate.Period)
	lctx := common.GetContextFromState(time.Now(), rate, expiration, count)
	return lctx, nil
}

func (store *Store) getCachekey(key string) string {
	buffer := strings.Builder{}
	buffer.WriteString(store.Prefix)
	buffer.WriteString(":")
	buffer.WriteString(key)
	return buffer.String()
}
