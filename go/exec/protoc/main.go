package main

import (
	"encoding/json"
	std "exec/go/exec/protoc/protocbuf"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	test := &std.Student{
		Name:   "geektutu",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	newData := &std.Student{}

	err = proto.Unmarshal(data, newData)
	fmt.Println(newData)
	if err != nil {
		fmt.Println(err)
	}

	da, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(da))
}
