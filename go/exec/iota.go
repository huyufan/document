package main

import "fmt"

const (
	one int = iota
	two
	three
	four
)

func main() {
	fmt.Println(two)
	var i = 1

	var j = 2

	a, b := j, i

	fmt.Println(a, b)

}
