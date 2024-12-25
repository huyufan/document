package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

var (
	host       string
	localPort  int
	remotePort int
)

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "remote server ip")
	flag.IntVar(&localPort, "l", 8080, "the local port")
	flag.IntVar(&remotePort, "r", 3333, "remote server port")
}

type server struct {
	conn  net.Conn
	read  chan []byte
	write chan []byte
	exit  chan error
	// 重连通道
	reConn chan bool
}

func (s *server) Read(ctx context.Context) {
	_ = s.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data := make([]byte, 1024)
			n, err := s.conn.Read(data)
			if err != nil && err != io.EOF {
				if strings.Contains(err.Error(), "timeout") {
					_ = s.conn.SetReadDeadline(time.Now().Add(time.Second * 3))
					s.conn.Write([]byte("pi"))
					continue
				}
				fmt.Println("从server读取数据失败, ", err.Error())
				s.exit <- err
				return
			}

			if data[0] == 'p' && data[1] == 'i' {
				fmt.Println("client收到心跳包")
				continue
			}
			s.read <- data[:n]
		}
	}
}

func (s *server) Write(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-s.write:
			_, err := s.conn.Write(data)
			if err != nil && err != io.EOF {
				s.exit <- err
				return
			}
		}
	}
}

type local struct {
	conn  net.Conn
	read  chan []byte
	write chan []byte
	exit  chan error
}

func (l *local) Read(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data := make([]byte, 1024)
			n, err := l.conn.Read(data)
			if err != nil {
				l.exit <- err
				return
			}
			l.read <- data[:n]
		}
	}
}

func (l *local) Write(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-l.write:
			_, err := l.conn.Write(data)
			if err != nil {
				l.exit <- err
				return
			}
		}
	}
}

func main() {
	flag.Parse()
	target := net.JoinHostPort(host, fmt.Sprintf("%d", host))
	for {
		serverConn, err := net.Dial("tcp", target)
		if err != nil {
			panic(err)
		}
		fmt.Printf("已连接server: %s \n", serverConn.RemoteAddr())
		server := &server{
			conn:   serverConn,
			read:   make(chan []byte),
			write:  make(chan []byte),
			exit:   make(chan error),
			reConn: make(chan bool),
		}
		go handle(server)
		<-server.reConn
	}
}

func handle(server *server) {
	ctx, cancel := context.WithCancel(context.Background())
	go server.Read(ctx)
	go server.Write(ctx)

	localConn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", localPort))
	if err != nil {
		panic(err)
	}
	local := &local{
		conn:  localConn,
		read:  make(chan []byte),
		write: make(chan []byte),
		exit:  make(chan error),
	}
	go local.Read(ctx)
	go local.Write(ctx)
	defer func() {
		_ = server.conn.Close()
		_ = local.conn.Close()
		server.reConn <- true
	}()

	for {
		select {
		case data := <-server.read:
			local.write <- data
		case data := <-server.read:
			server.write <- data
		case err := <-server.exit:
			fmt.Printf("server have err: %s", err.Error())
			cancel()
			return
		case err := <-local.exit:
			fmt.Printf("local have err: %s", err.Error())
			cancel()
			return
		}
	}

}
