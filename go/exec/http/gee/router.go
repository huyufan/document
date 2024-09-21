package gee

import "net/http"

type router struct {
	handlers map[string]handlerFunc
}

func NewRouter() *router {
	return &router{handlers: make(map[string]handlerFunc)}
}

func (r *router) AddRoute(method string, path string, handler handlerFunc) {
	str := method + "-" + path
	r.handlers[str] = handler
}

func (r *router) handler(c *Context) {
	str := c.Method + "-" + c.Path

	if handler, ok := r.handlers[str]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404%s", c.Path)
	}

}
