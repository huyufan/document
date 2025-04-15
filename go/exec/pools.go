package main

import (
	"sync"
	"time"
)

var pl = sync.Pool{
	New: func() interface{} {
		t := time.NewTimer(time.Second)
		// t.Stop()
		return t
	},
}

func main() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		ch <- 100
	}()

	select {
	case i := <-ch:
		println("case1 recv: ", i)
	case i := <-ch:
		println("case2 recv: ", i)
		// default:
		// 	println("default case")
	}
	// ch := make(chan int, 1)
	// select {
	// case ch <- getVal(1):
	// 	println("recv: ", <-ch)
	// case ch <- getVal(2):
	// 	println("recv: ", <-ch)
	// }
}

func getVal(n int) int {
	println("getVal: ", n)
	return n
}
