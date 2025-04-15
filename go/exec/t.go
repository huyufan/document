package main

import "fmt"

func Swap[x int](a, b x) (x, x) {
	return b, a
}

func main() {
	x, y := Swap(10, 20)
	fmt.Println(x, y)
}
