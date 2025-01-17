package main

import "fmt"

//策略模式
type IStrategy interface {
	do(int, int) int
}

type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) do(a, b int) int {
	return a - b
}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}

func main() {
	colo := Operator{strategy: &add{}}

	value := colo.calculate(4, 5)

	fmt.Println(value)
}
