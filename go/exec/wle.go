package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//v := map[string][string]{"1"=>"2"}
	sd := map[string]string{"name": "dd", "val": "tt"}

	jst, _ := json.Marshal(sd)

	fmt.Println(string(jst))

	s1 := []int{1, 2, 3, 4, 5}

	fmt.Println(s1[:len(s1)])
}
