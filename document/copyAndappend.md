# copy 和 append 是常用的内建函数，它们有不同的用途

## 用处比较
- **copy**：用于将一个切片的数据复制到另一个切片中。
- **append** ：用于向切片中追加元素，并在必要时扩展底层数组的容量。

## 性能比较
- copy 直接将一个切片的元素复制到另一个切片。其复杂度为 O(n)，其中 n 是被复制的元素数量。
- copy 函数只在两个切片之间进行元素拷贝，不会涉及底层数组的重新分配。因此，当目标切片的容量足够时，copy 的性能通常优于 append，因为它不涉及内存分配。
- append 追加元素时，如果原切片的容量不足，Go 会为切片分配新的底层数组，并将数据复制到新数组。这个过程涉及内存分配和元素复制，因此当需要扩容时，append 的性能可能比 copy 慢。

### 示例 
``` go
package main

import (
    "fmt"
    "time"
)

func main() {
    src := make([]int, 10000000)
    destCopy := make([]int, len(src))

    // 测试 copy
    start := time.Now()
    copy(destCopy, src)
    fmt.Println("Time taken by copy:", time.Since(start))

    // 测试 append
    destAppend := make([]int, 0, len(src)) // 预留足够的容量，避免扩容
    start = time.Now()
    destAppend = append(destAppend, src...)
    fmt.Println("Time taken by append:", time.Since(start))
}
```
