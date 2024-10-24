package stdlib

import (
	"exec/go/exec/limit"
	"net/http"
)

type Option interface {
	apply(*Middleware)
}

type option func(*Middleware)

func (o option) apply(middleware *Middleware) {
	o(middleware)
}

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

func WithErrorHandler(handler ErrorHandler) option {
	return func(middleware *Middleware) {
		middleware.OnError = handler
	}
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {

	defer func() {
		if r := recover(); r != nil {

			w.Header().Add("Content-type", "application/json")

			w.Write([]byte(`{"meaasge":"No"}`))
		}
	}()

}

type LimitReachedHandler func(w http.ResponseWriter, r *http.Request)

func WithLimitReachedHandler(handler LimitReachedHandler) option {
	return func(m *Middleware) {
		m.OnLimitReached = handler
	}
}

func DefaultLimitRachedHandler(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "Limit exceeded", http.StatusTooManyRequests)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"mesage":"no"}`))
}

type KeyGetter func(r *http.Request) string

func WithKeyGetter(handler KeyGetter) option {
	return func(m *Middleware) {
		m.OnKeyGetter = handler
	}
}

func DefaultKeyGetter(limiter *limit.Limiter) func(r *http.Request) string {
	return func(r *http.Request) string {
		return limiter.GetIPKey(r)
	}
}

func WithExcludekey(handler func(string) bool) option {
	return func(m *Middleware) {
		m.Excludedkey = handler
	}
}
