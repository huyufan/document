package main

import "fmt"

func outer(i int) func(y int) int {
	return func(y int) int {
		return i + y
	}
}

//不能嵌套具名函数：Go 不允许在函数内部定义具名函数。
//匿名函数可以嵌套：Go 允许在函数内部定义匿名函数，并且可以立即调用或将其作为值使用。

func main() {
	anon := func(x int) {
		fmt.Println(x)
	}
	anon(4)

	f := outer(10)
	fmt.Println(f(200))

	fmt.Println(test())

	for _, v := range test() {
		v()
	}

}

func test() []func() {
	var s []func()

	for i := 2; i < 5; i++ {
		s = append(s, func() {
			fmt.Println(&i, i)
		})
	}
	return s
}
