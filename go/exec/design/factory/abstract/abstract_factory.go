package main

import "fmt"

type ProductA interface {
	UseA()
}

type ProductA1 struct {
}

func (p *ProductA1) UseA() {
	fmt.Println("usea1")
}

type ProductA2 struct {
}

func (p *ProductA2) UseA() {
	fmt.Println("usea2")
}

type ProductB interface {
	UseB()
}

type ProductB1 struct {
}

func (p *ProductB1) UseB() {
	fmt.Println("useb1")
}

type ProductB2 struct {
}

func (p *ProductB2) UseB() {
	fmt.Println("useb2")
}

type abstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type factory1 struct{}

func (f *factory1) CreateProductA() ProductA {
	return &ProductA1{}
}

func (f *factory1) CreateProductB() ProductB {
	return &ProductB1{}
}

type factory2 struct{}

func (f *factory2) CreateProductA() ProductA {
	return &ProductA2{}
}

func (f *factory2) CreateProductB() ProductB {
	return &ProductB2{}
}
