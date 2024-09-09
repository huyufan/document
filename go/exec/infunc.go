package main

import "fmt"

type Serverer interface {
	Show(s *Servers)
}

type Servers struct {
	Name string
	Port int
}

type Options func(*Servers)

func (o Options) Show(s *Servers) {
	o(s)
}

func WithNames(name string) Serverer {
	return Options(func(s *Servers) {
		s.Name = name
	})
}

func NewServers(opt ...Serverer) *Servers {
	ser := &Servers{
		Name: "defualt",
		Port: 3306,
	}
	for _, o := range opt {
		o.Show(ser)
	}
	return ser
}
func main() {
	s := NewServers(WithNames("huyufan"))
	fmt.Println(s)
}
