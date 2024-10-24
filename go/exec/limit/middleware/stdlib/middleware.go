package stdlib

import (
	"exec/go/exec/limit"
	"net/http"
	"strconv"
)

type Middleware struct {
	Limiter        *limit.Limiter
	OnError        ErrorHandler
	OnLimitReached LimitReachedHandler
	OnKeyGetter    KeyGetter
	Excludedkey    func(string) bool
}

func NewMiddleware(limiter *limit.Limiter, options ...option) *Middleware {
	middleware := &Middleware{
		Limiter:        limiter,
		OnError:        DefaultErrorHandler,
		OnLimitReached: DefaultLimitRachedHandler,
		OnKeyGetter:    DefaultKeyGetter(limiter),
		Excludedkey:    nil,
	}

	for _, option := range options {
		option.apply(middleware)
	}
	return middleware
}

func (middleware *Middleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := middleware.OnKeyGetter(r)
		if middleware.Excludedkey != nil && middleware.Excludedkey(key) {
			h.ServeHTTP(w, r)
			return
		}
		context, err := middleware.Limiter.Get(r.Context(), key)
		if err != nil {
			middleware.OnError(w, r, err)
			return
		}
		w.Header().Add("X-RateLimit-Limitv", strconv.FormatInt(context.Limit, 10))
		w.Header().Add("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		w.Header().Add("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))
		if context.Reached {
			middleware.OnLimitReached(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})

}
