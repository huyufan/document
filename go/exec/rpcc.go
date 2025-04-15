package main

import (
	"fmt"
	"net/rpc"
)

type Param struct {
	Width, Height int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	ret := 0
	err2 := conn.Call("Rect.Area", Param{10, 20}, &ret)
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println(ret)
}
