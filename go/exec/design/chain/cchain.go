package main

import "fmt"

type Filter interface {
	Process(content string) bool
	SetNext(handler Filter)
}

// 广告过滤器
type AdFilter struct {
	next Filter
}

func (a *AdFilter) Process(content string) bool {
	if containsAd(content) {
		fmt.Println("广告内容检测到，拒绝请求")
		return false // 拒绝请求
	}
	if a.next != nil {
		return a.next.Process(content) // 传递给下一个过滤器
	}
	return true
}

func (a *AdFilter) SetNext(next Filter) {
	a.next = next
}

// 色情内容过滤器
type PornFilter struct {
	next Filter
}

func (p *PornFilter) Process(content string) bool {
	if containsPorn(content) {
		fmt.Println("色情内容检测到，拒绝请求")
		return false // 拒绝请求
	}
	if p.next != nil {
		return p.next.Process(content) // 传递给下一个过滤器
	}
	return true
}

func (p *PornFilter) SetNext(next Filter) {
	p.next = next
}

// 简单的广告检测函数
func containsAd(content string) bool {
	return content == "广告"
}

// 简单的色情内容检测函数
func containsPorn(content string) bool {
	return content == "色情"
}

func main() {
	// 创建过滤器链
	adFilter := &AdFilter{}
	pornFilter := &PornFilter{}

	// 设置责任链顺序
	adFilter.SetNext(pornFilter)

	// 测试内容
	content1 := "广告"
	content2 := "正常内容"
	content3 := "色情"
	content4 := "正常内容，无广告"

	// 测试广告过滤
	fmt.Println("测试1：", adFilter.Process(content1)) // 应该拒绝请求，广告检测到
	fmt.Println("测试2：", adFilter.Process(content2)) // 应该通过过滤器链，正常内容
	fmt.Println("测试3：", adFilter.Process(content3)) // 应该拒绝请求，色情内容检测到
	fmt.Println("测试4：", adFilter.Process(content4)) // 应该通过过滤器链，正常内容
}
