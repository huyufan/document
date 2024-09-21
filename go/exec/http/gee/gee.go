package gee

import (
	"fmt"
	"net/http"
)

type handlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (engine *Engine) AddRoute(method string, path string, handler handlerFunc) {
	engine.router.AddRoute(method, path, handler)
	fmt.Println(engine.router)
}

func (engine *Engine) GET(path string, handler handlerFunc) {
	engine.AddRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler handlerFunc) {
	engine.AddRoute("POST", path, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handler(c)
}
