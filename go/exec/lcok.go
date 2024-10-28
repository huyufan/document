package main

import (
	"fmt"
	"sync"
	"time"
)

type abs[T any] struct {
	s T
	b int
}

func nTest[T any](v T) abs[T] {
	return abs[T]{s: v}
}

var lock sync.RWMutex

func main() {
	dd := nTest(23)
	fmt.Println(dd.s)
	// go Rlock(1)
	// go Wlock()
	// go Rlock(2)
	// time.Sleep(20 * time.Second)
}

func Rlock(i int) {
	if i == 2 {
		time.Sleep(100 * time.Microsecond)
	}
	lock.RLock()
	defer lock.RUnlock()
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())
}

func Wlock() {
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("huyufan", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("huyufan", time.Now())
}
