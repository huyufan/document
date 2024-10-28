package main

import (
	"fmt"
)

type a int

type b = int

func (c a) add(s int) {
	fmt.Print(int(c) + s)
}

type Operation func(a int, b int) int

type callback func(a int)

func Check(a int, b int, op Operation) {
	op(a, b)
}

func Add(a int, b int) int {
	fmt.Println(a + b)
	return a + b
}

func Substrin(a int, b int) int {
	fmt.Println(a - b)
	return a - b
}

func sd(name string) Operation {
	if name == "add" {
		return func(a, b int) int {
			fmt.Println(a + b)
			return a + b
		}
	} else {
		return func(a, b int) int {
			fmt.Println(a - b)
			return a - b
		}
	}
}

func doaction(name string, call callback) {
	fmt.Println(name)
	d := 10
	call(d)
}

func main() {
	var s a = 12
	s.add(4)

	// Check(40, 5, Add)
	// Check(40, 5, Substrin)

	// tt := sd("add")
	// tt(3, 4)

	// doaction("huyufan", func(a int) {
	// 	fmt.Print("cc")
	// 	fmt.Println("dd")
	// 	fmt.Println(a)
	// })

}
