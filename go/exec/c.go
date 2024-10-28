package main

import (
	"fmt"
	"net/http"
)

type MyHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type LoggingMiddleware struct {
	next MyHandler
}

func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging Middleware1:")
	l.next.ServeHTTP(w, r)
	fmt.Println("Logging Middleware2:")
}

type AuthMiddleware struct {
	next MyHandler
}

func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authing Middleware3:")
	a.next.ServeHTTP(w, r)
	fmt.Println("Authing Middleware4:")
}

// 最终的业务处理器，实现 MyHandler 接口
type FinalHandler struct{}

// 实现 MyHandler 接口的方法
func (f *FinalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Final Handler: 处理业务逻辑")
	fmt.Fprintln(w, "业务处理完成，Hello, World!")
}

func main() {
	final := &FinalHandler{}

	auth := &AuthMiddleware{next: final}

	login := &LoggingMiddleware{next: auth}

	http.Handle("/", login)

	http.ListenAndServe(":8878", nil)

}
