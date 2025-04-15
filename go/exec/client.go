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
	conn   net.Conn
	read   chan []byte
	write  chan []byte
	exit   chan error
	reConn chan bool
}

func (s *server) Read(ctx context.Context) {
	_ = s.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data := make([]byte, 10240)
			n, err := s.conn.Read(data)
			if err != nil && err != io.EOF {
				if strings.Contains(err.Error(), "timeout") {
					_ = s.conn.SetReadDeadline(time.Now().Add(time.Second * 3))
					s.conn.Write([]byte("pi"))
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
