package main

import "fmt"

type server struct {
	name string
	port int
}

type Options func(*server)

func WithName(name string) Options {
	return func(o *server) {
		o.name = name
	}
}

func Withport(port int) Options {
	return func(o *server) {
		o.port = port
	}
}

func Newop(opt ...Options) *server {
	op := &server{}

	for _, o := range opt {
		o(op)
	}
	return op
}

func main() {
	sd := Newop(WithName("huyufan"))
	fmt.Println(sd)
}
