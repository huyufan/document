package main

import (
	"bufio"
	"fmt"
	"os"
)

// 缓冲读取（Buffered Reading）和非缓冲读取（Unbuffered Reading）的区别在于，
// 缓冲读取会在底层 I/O 操作与应用程序之间加上一个缓冲区，用来减少系统调用的次数，提升读取性能。缓冲读取的原理是将一部分数据提前读入内存缓冲区，
// 后续读取操作直接从内存中取数据，而不必频繁地与底层数据源（如文件或网络）交互。
func main() {
	file, err := os.Open("func.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buf := make([]byte, 10)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Read error:", err)
			return
		}
		fmt.Println(n)
		//fmt.Print(string(buf[:n]))
	}
}
