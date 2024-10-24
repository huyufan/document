package limit

import (
	"context"
	"time"
)

type Store interface {
	Get(ctx context.Context, key string, rate Rate) (Context, error)
	Peek(ctx context.Context, key string, rate Rate) (Context, error)
	Reset(ctx context.Context, key string, rate Rate) (Context, error)
	Increment(ctx context.Context, key string, count int64, rate Rate) (Context, error)
}

type StoreOptions struct {
	Prefix          string
	MaxRetry        int
	CleanUpInterval time.Duration
}
