package main

import (
	"encoding/json"
	"exec/go/exec/alice"
	"fmt"
	"net/http"
	"strconv"
)

type Results struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func Recover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("huyufan")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("huyufan")
				fmt.Println(err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("huyufanddd")
		h.ServeHTTP(w, r)
	})
}

func Routes() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//ba := r.PathValue("back")
		bk := r.PathValue("bk")
		sd := []int{1, 2, 4}
		key, _ := strconv.Atoi(bk)
		_ = sd[key]

		w.Header().Set("Content-Type", "application/json")
		result := Results{Code: 200, Message: []int{1, 2, 3, 4}}

		_ = json.NewEncoder(w).Encode(result)
	}

}

func main() {
	alices := alice.New(Recover, Logging)

	http.Handle("GET /b/{back}/0/{bk}", alices.Then(Routes()))

	http.Handle("GET /c/{back}/0/{bk}", Recover(Logging(Routes())))

	http.ListenAndServe(":8888", nil)
}
