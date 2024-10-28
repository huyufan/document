package main

import (
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func te(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 中间件逻辑，例如打印请求信息
		fmt.Println("Request received:", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
		fmt.Println("huyfan")
	})
}

func re(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 中间件逻辑，例如打印请求信息
		fmt.Println("re:", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
		fmt.Println("erhuyfan")
	})
}

func main() {

	http.Handle("GET /b/{back}/0/{bk}", te(http.HandlerFunc(indexs)))

	http.ListenAndServe(":8888", nil)
}

func indexs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;chartset=utf-8")
	_, err := w.Write([]byte(`{"meaasge":"ok}`))
	if err != nil {
		log.Fatalln(err)
	}
}

// http.HandleFunc("GET /b/{back}/0/{bk}", func(w http.ResponseWriter, r *http.Request) {
// 	ba := r.PathValue("back")
// 	bk := r.PathValue("bk")
// 	fmt.Println(ba)
// 	fmt.Println(bk)
// 	w.Header().Set("Content-Type", "application/json")
// 	result := Result{Code: 200, Message: []int{1, 2, 3, 4}}

// 	_ = json.NewEncoder(w).Encode(result)
// })
