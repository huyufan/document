# 单元测试
- 在 Go 语言中，单元测试通过使用标准库中的 testing 包来编写和执行。测试文件通常以 _test.go 结尾，并且测试函数的命名必须以 Test 开头。每个测试函数接受一个 *testing.T 参数，用于记录测试的状态和输出。
```  go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

```
## 接下来，我们为这个函数编写单元测试，测试文件应该命名为 math_test.go：
``` go
// math_test.go
package math

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

```
## 运行测试 go test

### 通过
``` shell
ok  	package_name	0.001s
```

### 不通过
``` shell
--- FAIL: TestAdd (0.00s)
    math_test.go:10: Add(2, 3) = 4; want 5
FAIL
exit status 1
FAIL	package_name	0.001s
```
## 测试覆盖率
``` shell 
go test -cover

```
### 输出将显示测试覆盖率的百分比：
``` shell 
ok  	package_name	0.001s	coverage: 100.0% of statements

```


# 基准测试
- 除了单元测试，Go 还支持基准测试，用于衡量代码的性能。基准测试函数以 Benchmark 开头，接收一个 *testing.B 参数：
``` go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

```
## 运行基准测试：
``` shell 
go test -bench=.
```

##  测试初始化代码
- 如果需要在测试前执行一些初始化代码，可以使用 TestMain 函数：

``` go
func TestMain(m *testing.M) {
    // 设置测试环境
    code := m.Run() // 运行测试
    // 清理操作
    os.Exit(code)
}

```