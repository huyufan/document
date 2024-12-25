package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	localPort  int
	remotePort int
)

func init() {
	flag.IntVar(&localPort, "l", 5200, "the user link port")
	flag.IntVar(&remotePort, "r", 3333, "client listen port")
}

type client struct {
	conn net.Conn
	// 数据传输通道
	read  chan []byte
	write chan []byte
	// 异常退出通道
	exit chan error
	// 重连通道
	reConn chan bool
}

func (c *client) Read(ctx context.Context) {
	_ = c.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data := make([]byte, 1024)
			n, err := c.conn.Read(data)
			if err != nil && err != io.EOF {
				if strings.Contains(err.Error(), "timeout") {
					_ = c.conn.SetReadDeadline(time.Now().Add(time.Second * 3))
					c.conn.Write([]byte("pi"))
					continue
				}
				fmt.Println("读取出现错误...")
				c.exit <- err
				return
			}
			// 收到心跳包,则跳过
			if data[0] == 'p' && data[1] == 'i' {
				fmt.Println("server收到心跳包")
				continue
			}

			c.read <- data[:n]
		}
	}
}

func (c *client) Write(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-c.write:
			_, err := c.conn.Write(data)
			if err != nil && err != io.EOF {
				c.exit <- err
				return
			}
		}
	}
}

type user struct {
	conn  net.Conn
	read  chan []byte
	write chan []byte
	exit  chan error
}

func (u *user) Read(ctx context.Context) {
	_ = u.conn.SetReadDeadline(time.Now().Add(time.Second * 200))
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data := make([]byte, 1024)
			n, err := u.conn.Read(data)
			if err != nil && err != io.EOF {
				u.exit <- err
				return
			}
			u.read <- data[:n]
		}
	}
}

func (u *user) Write(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-u.write:
			_, err := u.conn.Write(data)
			if err != nil && err != io.EOF {
				u.exit <- err
				return
			}
		}
	}
}

func main() {
	flag.Parse()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	clientListener, err := net.Listen("tcp", strconv.Itoa(remotePort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("监听:%d端口, 等待client连接... \n", remotePort)

	userListener, err := net.Listen("tcp", strconv.Itoa(localPort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("监听:%d端口, 等待user连接.... \n", localPort)

	for {
		clientConn, err := clientListener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("有Client连接: %s \n", clientConn.RemoteAddr())
		client := &client{
			conn:   clientConn,
			read:   make(chan []byte),
			write:  make(chan []byte),
			exit:   make(chan error),
			reConn: make(chan bool),
		}
		userConnChan := make(chan net.Conn)
		go AcceptUserConn(userListener, userConnChan)
		go HandleClient(client, userConnChan)

	}
}

func AcceptUserConn(userListener net.Listener, connChan chan net.Conn) {
	userConn, err := userListener.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Printf("user connect: %s \n", userConn.RemoteAddr())
	connChan <- userConn
}

func HandleClient(client *client, userConnChan chan net.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	go client.Read(ctx)
	go client.Write(ctx)
	user := &user{
		read:  make(chan []byte),
		write: make(chan []byte),
		exit:  make(chan error),
	}

	defer func() {
		_ = client.conn.Close()
		_ = user.conn.Close()
		client.reConn <- true
	}()

	for {
		select {
		case userConn := <-userConnChan:
			user.conn = userConn
			go handle(ctx, client, user)
		case err := <-client.exit:
			fmt.Println("client出现错误, 关闭连接", err.Error())
			cancel()
			return
		case err := <-user.exit:
			fmt.Println("user出现错误，关闭连接", err.Error())
			cancel()
			return
		}
	}
}

func handle(ctx context.Context, client *client, user *user) {
	go user.Read(ctx)
	go user.Write(ctx)

	for {
		select {
		case userRecv := <-user.read:
			client.write <- userRecv
		case clientRecv := <-client.write:
			user.read <- clientRecv
		case <-ctx.Done():
			return
		}
	}
}
