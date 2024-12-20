package main

import "fmt"

type Girl interface {
	weight()
}

type FatGril struct{}

func (FatGril) weight() {
	fmt.Println("80kg")
}

type ThinGirl struct {
}

func (ThinGirl) weight() {
	fmt.Println("40kg")
}

type GirlFactory struct {
}

func (*GirlFactory) CreateGirl(like string) Girl {
	switch like {
	case "fat":
		return &FatGril{}
	case "thin":
		return &ThinGirl{}
	}
	return nil
}
