package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func s() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

func t() (a int) {
	defer func() {
		a++
	}()
	return 1
}

func GoA() {
	defer (func() {
		if err := recover(); err != nil {
			fmt.Println("panic:" + fmt.Sprintf("%s", err))
		}
	})()

	go GoB()
}

func GoB() {
	panic("error")
}
func main() {

	jsonStr := `{"number":1234567}`
	result := make(map[string]interface{})
	d := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	d.UseNumber()
	err := d.Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	// p4 := struct {
	// 	Name string
	// 	Age  int
	// }{Name: "hyf", Age: 33}

	// fmt.Println(p4)
	// s1 := []int{1, 2, 3, 4, 5}
	// fmt.Println(append(s1[:1], s1[3:5]...))

	// funcSml := func(c int) {
	// 	fmt.Println(c)
	// }
	// for _, v := range te() {
	// 	v()
	// }
	//funcSml(6)
}

func te() []func() {
	var fun []func()

	for i := 1; i < 5; i++ {
		fun = append(fun, func() { fmt.Println(i) })
	}
	return fun
}
