package singleflight

import "sync"

type Call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*Call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*Call)
	}
	if val, ok := g.m[key]; ok {
		g.mu.Unlock()
		val.wg.Wait()
		return val.val, nil
	}
	c := new(Call)
	g.m[key] = c
	c.wg.Add(1)
	c.val, c.err = fn()
	c.wg.Done()
	delete(g.m, key)
	g.mu.Unlock()
	return c.val, c.err
}
