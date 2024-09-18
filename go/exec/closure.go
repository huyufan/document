package main

import (
	"fmt"
	"sync"
)

func main() {
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
}
