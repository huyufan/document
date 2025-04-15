package main

import (
	"fmt"
	"io"
)

// 函数赋值
func add(a, b int) int {
	return a + b
}

// 函数返回值
func pow(a int) func(int) int {
	return func(l int) int {
		result := 1
		for i := 0; i < l; i++ {
			result *= a
		}
		return result
	}
}

// 函数参数传值
func filter(a []int, fn func(int) bool) (result []int) {
	for _, v := range a {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func generateInteger() func() int {
	ch := make(chan int)
	count := 0
	go func() {
		for {
			ch <- count
			count++
		}
	}()

	return func() int {
		return <-ch
	}
}

func main() {
	//变量复制
	fn := add
	result := fn(3, 4)

	fmt.Println(result)

	p := pow(2)

	res := p(3)
	fmt.Println(res)

	a := []int{1, 2, 3, 4, 5}

	ad := filter(a, func(i int) bool {
		return i&1 == 1
	})

	fmt.Println(ad)

	ads := filter(a, func(i int) bool {
		return i&1 == 0
	})

	fmt.Println(ads)

	generate := generateInteger()

	fmt.Println(generate())

	fmt.Println(generate())
	fmt.Println(generate())

	type Person struct {
		firstName string
		lastName  string
		fullName  string
		age       int
	}
	var getFullName = func(in *Person) string {
		in.fullName = in.firstName + in.lastName
		return in.fullName
	}

	john := Person{
		"john", "doe", "", 30,
	}

	fmt.Println(getFullName(&john))
	fmt.Println(john)

	var w io.Writer
	fmt.Fprint(w, "qwqw")
}
