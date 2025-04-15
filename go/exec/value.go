package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var rdb *redis.Client
var ctx = context.Background()

// 假设这个是我们要执行的耗时操作
func fetchDataFromDB() (interface{}, error) {
	// 模拟一个耗时操作
	time.Sleep(2 * time.Second)
	return "Fetched Data from DB", nil
}

// 获取数据的函数
func getData(key string) (interface{}, error) {
	// 先检查 Redis 是否有缓存
	cachedValue, err := rdb.Get(ctx, key).Result()
	if err == nil {
		// 如果 Redis 中有数据，直接返回
		return cachedValue, nil
	}

	// 如果 Redis 中没有数据，则执行操作
	// 为了防止多个服务器同时执行相同的操作，使用 Redis 锁
	lockKey := key + ":lock"

	// 尝试获取锁，设置超时防止死锁
	ok, err := rdb.SetNX(ctx, lockKey, 1, 10*time.Second).Result()
	if err != nil {
		return nil, err
	}

	if ok {
		// 锁定成功，开始执行操作
		data, err := fetchDataFromDB()
		if err != nil {
			return nil, err
		}

		// 将结果缓存到 Redis 中，设置过期时间
		rdb.Set(ctx, key, data, 30*time.Second)

		// 释放锁
		rdb.Del(ctx, lockKey)

		return data, nil
	}

	// 锁定失败，说明其他服务器正在执行操作，等待结果
	// 使用一个类似于 `SingleFlight` 的机制来等待结果
	for {
		cachedValue, err := rdb.Get(ctx, key).Result()
		if err == nil {
			return cachedValue, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 连接 Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
		DB:   0,                // 使用默认数据库
	})

	// 测试多个 goroutine 并发调用相同操作
	for i := 0; i < 5; i++ {
		go func(i int) {
			key := "unique-key"
			data, err := getData(key)
			if err != nil {
				log.Printf("Error in goroutine %d: %v", i, err)
				return
			}
			fmt.Printf("Goroutine %d received data: %v\n", i, data)
		}(i)
	}

	// 等待所有 goroutine 执行完
	time.Sleep(5 * time.Second)
}
