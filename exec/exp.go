package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "huyufan")
	})

	http.HandleFunc("/huyu", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "huyu")
		//os.Exit(0)
	})

	http.ListenAndServe(":8888", nil)
}
