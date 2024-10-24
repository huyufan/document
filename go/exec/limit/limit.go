package limit

import "context"

type Context struct {
	Limit     int64
	Remaining int64
	Reset     int64
	Reached   bool
}

type Limiter struct {
	store   Store
	rate    Rate
	Options Options
}

func New(store Store, rate Rate, options ...Option) *Limiter {
	opt := Options{
		IPv4Mask:           DefaultIPv4Mask,
		IPv6Mask:           DefaultIPv6Mask,
		TrustForwardHeader: false,
	}

	for _, option := range options {
		option(&opt)
	}

	return &Limiter{
		store:   store,
		rate:    rate,
		Options: opt,
	}

}

func (limiter *Limiter) Get(ctx context.Context, key string) (Context, error) {
	return limiter.store.Get(ctx, key, limiter.rate)
}

func (limiter *Limiter) Peek(ctx context.Context, key string) (Context, error) {
	return limiter.store.Peek(ctx, key, limiter.rate)
}

func (limiter *Limiter) Reset(ctx context.Context, key string) (Context, error) {
	return limiter.store.Reset(ctx, key, limiter.rate)
}

func (limiter *Limiter) Increment(ctx context.Context, key string, count int64) (Context, error) {
	return limiter.store.Increment(ctx, key, count, limiter.rate)
}
