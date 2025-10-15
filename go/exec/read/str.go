package main

import (
	"fmt"
	"strings"
)

func main() {
	query := "sd=12&d[]=2&d[]=3"

	for query != "" {
		var a string
		a, query, _ = strings.Cut(query, "&")

		fmt.Println(a)
		fmt.Println(query)

		if a == "" {
			continue
		}

		// time.Sleep(4 * time.Second)

	}

}
