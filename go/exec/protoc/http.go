package main

import (
	std "exec/go/exec/protoc/protocbuf"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8888", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	std := &std.Student{
		Name:   "huyufan",
		Male:   true,
		Scores: []int32{2, 5, 10},
	}
	data, _ := proto.Marshal(std)
	w.Header().Set("Content-Type", "application/x-protobuf")

	html := "<html><body><p>我是睡</p></body></html>"
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))

}
