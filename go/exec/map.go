package main

import "fmt"

func main() {
	var ff map[int]int

	ff = make(map[int]int)

	ff[0] = 2

	c, ok := ff[1]

	fmt.Println(c, ok)
}
