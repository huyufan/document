package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

// sync.Pool 是并发安全的，可以在多协程环境下安全使用。
// 清理缓冲区：在将缓冲区放回池中之前，务必调用 buf.Reset() 清空缓冲区，防止数据污染。
type Student struct {
	Name string
	Age  int8
	Sex  int8
}

var buf, _ = json.Marshal(Student{Name: "huyufan", Age: 33, Sex: 1})
var pool = sync.Pool{
	New: func() any {
		return new(Student)
	},
}

func (s *Student) Reset() {
	s.Name = ""
	s.Age = 0
	s.Sex = 1
}

func main() {

	stu := pool.Get().(*Student)

	json.Unmarshal(buf, stu)
	fmt.Println(stu.Name)
	stu.Reset()
	pool.Put(stu)
}
