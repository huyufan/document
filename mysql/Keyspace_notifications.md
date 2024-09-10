# Keyspace Notifications 监听键的过期事件

## 详细实现步骤：

### 开启 Redis 的 Keyspace Notifications
- Redis 默认并不会启用 Keyspace Notifications，你需要手动开启。可以在 Redis 配置文件 redis.conf 中修改，或者通过 Redis 命令行开启。

### 通过命令开启：
``` shell
redis-cli config set notify-keyspace-events Ex  
```
- 其中 Ex 表示监听过期事件（E 为过期事件，x 为 key 过期的操作）。

### Go 实现监听器：
- 使用 SUBSCRIBE 命令订阅 Redis 的 Keyspace Notifications。
- 当监听到过期事件时，处理后续逻辑。

``` go
package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
    "context"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    // 订阅 key 失效事件
    pubsub := rdb.PSubscribe(ctx, "__keyevent@0__:expired")
    
    // 监听通道消息
    go func() {
        for msg := range pubsub.Channel() {
            fmt.Println("Received key expired event:", msg.Payload)
            // 处理后续逻辑
            handleKeyExpired(msg.Payload)
        }
    }()
    
    // 模拟一个订单并设置过期时间
    rdb.Set(ctx, "order_12345", "order details", 10*time.Minute)
    
    // 阻塞主线程
    select {}
}

func handleKeyExpired(key string) {
    fmt.Println("Key expired:", key)
    // 在这里可以处理订单后续逻辑
}

```

- **__keyevent@0__:expired：** Redis 的 Keyspace Notifications 频道格式，其中 @0 表示数据库 0，expired 表示监听的是过期事件。