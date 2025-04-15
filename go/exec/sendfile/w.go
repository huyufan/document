package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 监听 TCP 连接
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("服务器启动，等待客户端连接...")

	// 接受客户端连接
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 打开文件用于写入
	outFile, err := os.Create("r.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// 使用 io.Copy 将数据从连接读取并写入文件
	_, err = io.Copy(outFile, conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文件接收完成！")
}
