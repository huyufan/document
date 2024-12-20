package main

import "fmt"

type Customer interface {
	Update()
}

//观察者
type One struct{}

func (o *One) Update() {
	fmt.Println("我是第一个")
}

type Two struct{}

func (t *Two) Update() {
	fmt.Println("我是第二个")
}

// 被观察者

type NewOffice struct {
	customer []Customer
}

func (n *NewOffice) addCustomer(customer Customer) {
	n.customer = append(n.customer, customer)
}
func (n *NewOffice) newspaperCome() {
	n.notifyAllCustomer()
}

func (n *NewOffice) notifyAllCustomer() {
	for _, cus := range n.customer {
		cus.Update()
	}
}

func main() {
	a := &One{}
	b := &Two{}
	off := &NewOffice{}
	off.addCustomer(a)
	off.addCustomer(b)
	off.newspaperCome()
}
