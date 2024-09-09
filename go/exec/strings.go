package main

import (
	"fmt"
	"strings"
)

func main() {

	//len(s) 计算的是字符串 字节长度，也就是底层的 UTF-8 字节数。
	//len([]rune(s)) 计算的是字符串中的 字符（rune）数量，rune 表示一个 Unicode 字符。
	//字节长度是字符串存储的基础：在 Go 中，字符串是以字节序列存储的，因此拼接操作也是基于字节进行的
	var str strings.Builder
	s := "s我"
	str.Grow(len((s)))
	str.WriteString(s)
	fmt.Println(str.String())
}
