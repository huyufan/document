package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan struct{})
	done := make(chan struct{})

	go func() {
		ch1 <- struct{}{}
		close(ch1)
	}()

	go func() {
		for {
			select {
			case _, ok := <-ch1:
				if !ok {
					fmt.Println("Channel closed!")
					close(done) // 通知主协程退出
					return
				}
				fmt.Println("Received:", ok)
			}
		}
	}()

	<-done // 等待退出信号
	fmt.Println("over")
}
