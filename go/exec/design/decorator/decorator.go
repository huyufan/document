package main

import "fmt"

type Person interface {
	cost() int
	show()
}

type laowang struct {
}

func (*laowang) show() {
	fmt.Println("赤裸裸的老王。。。")
}

func (*laowang) cost() int {
	return 0
}

// 皮甲
type clothesDecorator struct {
	person Person
}

func (c *clothesDecorator) cost() int {
	return c.person.cost() + 20

}

func (c *clothesDecorator) show() {

}

func main() {
	lao := &laowang{}

	cloth := clothesDecorator{person: lao}

	sr := cloth.cost()

	fmt.Println(sr)
}
