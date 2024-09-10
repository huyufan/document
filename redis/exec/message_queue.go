package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// fmt.Println(client)

	// 启动生产者
	go func() {
		for i := 0; i < 10; i++ {
			client.LPush(ctx, "test", 55)
		}
	}()
	// // 启动消费者，使用阻塞 RPOP
	go func() {
		for {
			val, err := client.BRPop(ctx, 0, "test").Result() // 使用 BRPop 进行阻塞读取
			if err != nil {
				fmt.Println(12)
				fmt.Println(err)
				return
			}
			fmt.Println(val) // val[0] 是 key，val[1] 是实际的值
		}
	}()

	select {} // 阻塞主线程
}
