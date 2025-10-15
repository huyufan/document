package main

import (
	"bytes"
	"container/list"
	"fmt"
	"math"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

func put(buf *bytes.Buffer) {
	buf.Reset()
	bufPool.Put(buf)
}

func Abs(base, n int) int {
	m := n >> (base - 1)
	return (n + m) ^ m
}
func main() {
	s := Abs(64, -92222222222)
	fmt.Println(s)
	e := math.Sqrt(6)
	fmt.Println(e)
	// r := 5 ^ 1
	// c := map[string][]int{}
	// c["hyyf"] = []int{1}
	// c["ccc"] = []int{4}
	// c["hyyf"] = append(c["hyyf"], 4)
	// fmt.Println(c)
	// fmt.Println(r)

	list.Element

}
