package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false

}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	fmt.Println("无环链表:", hasCycle(node1)) // false

	// 创建有环链表：1 -> 2 -> 3 -> 2 (环)
	node3.Next = node1                    // 形成环
	fmt.Println("有环链表:", hasCycle(node1)) // true
}
