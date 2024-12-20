package main

import "sync"

type singleton struct{}

var (
	once sync.Once
	ins  *singleton
)

func NewSing() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
