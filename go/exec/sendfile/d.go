package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 打开文件
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 使用 io.Copy 将文件内容发送到连接
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文件已发送！")
}
