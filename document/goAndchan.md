# Go 协程与通道

## 协程 (Goroutine)

Go 协程（Goroutine）是与其他函数**同时运行的函数**，可以认为是轻量级的线程，由 Go 运行时进行管理。

### 协程的特点：
- **轻量级**：相比传统线程更为轻量，每个 Goroutine 只需少量内存开销。
- **并发执行**：协程可以并发运行，使得 Go 程序能够高效利用多核 CPU。
- **简单易用**：通过 `go` 关键字可以轻松启动一个 Goroutine。

### 示例代码：
```go
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printNumbers() // 启动一个新的协程
	fmt.Println("Main Goroutine running...")
	time.Sleep(6 * time.Second) // 保证主协程不退出
}
```


## 通道(Channel)

通道是协程之间通信的管道，用于传递数据，实现协程之间的同步与通信。

### 通道的特点：
- 线程安全：通道可以在多个协程之间传递数据，且无需显式的锁机制，避免了竞争条件。
- 阻塞机制：通道发送和接收数据时都是阻塞的，直到另一端准备好接收/发送数据为止。
- 方向性：  通道可以是双向的（可发送、接收），也可以是单向的（只发送或只接收）。

### 示例代码：
``` go
func main 

import "fmt"

func sendData(ch chan int){
	ch <- 10
}
func main(){
	ch := make(chan int)
	go sendData(ch)
	value := <-ch 
	fmt.Printf("value %d:",value)
}

```