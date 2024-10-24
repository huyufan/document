package alice

import (
	"fmt"
	"net/http"
)

type Constructor func(http.Handler) http.Handler

type Chain struct {
	constructors []Constructor
}

func New(constructor ...Constructor) Chain {
	return Chain{append(([]Constructor)(nil), constructor...)}
}

func (c Chain) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}
	fmt.Println(c.constructors)
	for i := range c.constructors {
		cc := c.constructors[len(c.constructors)-1-i]
		fmt.Println(cc)
		h = c.constructors[len(c.constructors)-1-i](h)
	}
	fmt.Println(h)
	return h
}

func (c Chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	if fn == nil {
		return c.Then(nil)
	}
	return c.Then(fn)
}

func (c Chain) Append(constructors ...Constructor) Chain {
	c.constructors = append(c.constructors, constructors...)
	return c
}

func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.constructors...)
}
