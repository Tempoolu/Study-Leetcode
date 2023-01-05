package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	fmt.Println(reverseList(list))
}

// 反转链表，简单，非常常见，可以做出来
func reverseList(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
