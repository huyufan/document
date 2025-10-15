package main

import (
	"fmt"
)

type strr struct {
	t string
}

type hc func(*strr)

type hcs []hc
type st struct {
	c int
}

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("goroutine A received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <-b
	fmt.Println("goroutine B received data: ", val)
	return
}

func a(ac ...hc) {
	b(ac)
}

func b(acs hcs) {
	for _, opt := range acs {
		opt(&strr{t: "1212"})
	}
	fmt.Println(acs)
}

func main() {
	hd := func(c *strr) {
		fmt.Println(c.t)
	}

	hs := func(c *strr) {
		fmt.Println(c.t)
	}

	a(hd, hs)

	// ch := make(chan int)
	// go goroutineA(ch)
	// go goroutineB(ch)
	// ch <- 3
	// ch <- 4
	// time.Sleep(time.Second)

	// dd := unsafe.Sizeof(struct{}{})
	// fmt.Println(dd)
	// de := []byte("我")
	// fmt.Println(de[0])
	// fmt.Println('0')
	// nq := new(st)

	// fmt.Println(nq)

	// var d = [256]uint8{'w': 3, 'c': 5}
	// fmt.Println(d)
	// //var asciiSpaces = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
	// cd := [6]uint{0: 1, 3: 2, 3, 4}

	// fmt.Println(cd)

	// str := "sdh,动平衡,drhr"
	// //fmt.Println(len(str))
	// n := 0
	// for i := 0; i < len(str); n++ {
	// 	fmt.Println(string(str[i]))
	// 	i++
	// }
	// //fmt.Println(n)
	// ss := strings.Split(str, ",")

	// fmt.Println(ss)

	// c := ^int(0)
	// fmt.Println(c)

	// l := list.New()
	// l.PushBack("A")
	// l.PushBack("B")
	// l.PushBack("C")
	// for e := l.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	// b1 := &strings.Builder{}

	// b1.WriteString("sss")

	// c2 := b1

	// c2.WriteString("sssww")
	// fmt.Printf("%+v", c2)
	// fmt.Printf("%+v", b1)
}
