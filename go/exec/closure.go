package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int = 1
	y := func() {
		x += 1

	}

	y()

	fmt.Println(x, y)
	var s []string = []string{"huyufan", "huyu", "fan"}
	var wg sync.WaitGroup

	wg.Add(len(s))
	for _, val := range s {
		go func(i string) {
			fmt.Println(i)
			wg.Done()
		}(val)
	}
	//注意由于匿名函数可以访问函数体外部的变量，而 for range 返回的 val 的值是引用的同一个内存地址的数据，所以匿名函数访问的函数体外部的 val 值是循环中最后输出的一个值
	wg.Wait()

	c, d := 1, 2

	//defer 调用会在当前函数执行结束前才被执行，这些调用被称为延迟调用，
	//defer 中使用匿名函数依然是一个闭包。
	defer func(a int) {
		fmt.Printf("c:%d,d:%d\n", a, d) // y 为闭包引用
	}(c) // 复制 x 的值

	c += 100
	d += 100
	fmt.Println(c, d)
}
