package main

import "fmt"

type Server struct {
	Name string
	Port int
}

type Option func(*Server)

func WithName(name string) Option {
	return func(s *Server) {
		s.Name = name
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}

}

func NewServer(opt ...Option) *Server {
	ser := &Server{
		Name: "default",
		Port: 3306,
	}

	for _, op := range opt {
		op(ser)
	}
	return ser
}

func main() {
	s := NewServer(WithName("huyufan"), WithPort(19604))
	fmt.Println(s)
}
