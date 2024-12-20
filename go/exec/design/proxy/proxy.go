package main

import "fmt"

type Seller interface {
	sell(name string)
}

// 火车站
type Station struct {
	stock int
}

func (s *Station) sell(name string) {
	if s.stock > 0 {
		s.stock--
		fmt.Println("代理点中：%s买了一张票,剩余：%d \n", name, s.stock)
	}
}

type StationProxy struct {
	station *Station
}

func (s *StationProxy) sell(name string) {
	if s.station.stock > 0 {
		s.station.sell(name)
	}
}

func main() {
	s := Station{stock: 100}

	proxy := &StationProxy{station: &s}

	proxy.sell("aohhas")
}
