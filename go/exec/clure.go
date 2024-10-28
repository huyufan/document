package main

import "fmt"

type Hd interface {
	ServeRequest(string)
}

type HdFunc func(string)

// HdFunc 是一个适配器，使函数类型实现 Hd 接口
func (h HdFunc) ServeRequest(s string) {
	h(s)
}

func L(h Hd) Hd {
	return HdFunc(func(s string) {
		fmt.Println("LoggingMiddleware: 记录日志", s)
		h.ServeRequest(s)
	})
}

func R(h Hd) Hd {
	return HdFunc(func(s string) {
		fmt.Println("rMiddleware: 记录日志", s)
		h.ServeRequest(s)
	})
}

type MyHd struct {
}

func (m MyHd) ServeRequest(s string) {
	fmt.Println("MyHandler: 处理请求", s)
}

func SR() Hd {
	return HdFunc(func(s string) {
		fmt.Println(s)
	})
}

func main() {

	//fin := &MyHd{}

	fin := SR()

	hand := R(L(fin))

	hand.ServeRequest("woshi")
}
