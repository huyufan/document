package main

import "fmt"

type Optioner interface {
	apply(s *servers)
}

type option func(*servers)

func (o option) apply(s *servers) {
	o(s)
}

type servers struct {
	Name string
	Port int
}

func WithNames(name string) Optioner {
	return option(func(s *servers) {
		s.Name = name
	})
}

func Newserv(opt ...Optioner) *servers {
	s := &servers{}

	for _, o := range opt {
		o.apply(s)
	}
	return s
}

func main() {
	cd := Newserv(WithNames("hf"))
	fmt.Println(cd)
}
