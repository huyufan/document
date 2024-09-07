# fmt.Println

## int指针 结构指针

### 示例代码：
``` go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Student struct {
	Name string
	Age  int8
	Sex  int8
}
func main() {
	var i *int = new(int)
	fmt.Println(i) //0xc00000a1e0

	stu := new(Student)
	fmt.Println(stu) //&{ 0 0}
}
```

### 原因
- 指针类型 (*int)：
当你打印 *int 类型的指针时，fmt.Println 默认输出指针的地址，因为它无法自动格式化指针指向的值（因为不知道指向的类型）。

-  结构体指针 (*Student)：
对于结构体类型的指针，fmt.Println 默认输出指针的地址以及指针指向的结构体的内容（以 &{...} 的格式），因为 Go 的 fmt 包有特定的处理规则来格式化结构体。
