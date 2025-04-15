package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	sdt = flag.Int("ccc", 1, "sdsdsd")
)

func main() {
	flag.Parse()
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(*sdt)
	ch := make(chan chan int)
	go func() {
		for {
			select {
			case c := <-ch:
				fmt.Println(c)
				go func(c chan int) {
					d := rand.Int()
					c <- d
				}(c)
			}
		}
	}()
	go func() {
		for i := 0; i < 3; i++ {
			chs := make(chan int)
			go func(s int, chh chan int) {
				ch <- chs
				message := <-chs
				fmt.Println(message)
			}(i, chs)
		}
	}()

	time.Sleep(2 * time.Second)
}
