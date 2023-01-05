package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head0 := &ListNode{Val: 1, Next: nil}
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	// head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	fmt.Println(removeNthFromEnd(head1, 2))
}

// 删除链表距离尾节点的第n个节点，中等，自己做出来
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/submissions/
// 画图，最重要是找边界
// 应该使用虚拟头，这样可以不用判断那么多边界
// 还是用了快慢指针

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	dummyHead := &ListNode{Next: head}
	fast := head
	prev := dummyHead
	var slow *ListNode
	count := 0

	for fast != nil {

		if count >= n {
			prev = slow
			slow = slow.Next
		}
		count++
		fast = fast.Next
	}
	if slow != nil {
		prev.Next = slow.Next
	}

	if count == n {
		head = head.Next
	}

	printList(head)
	return head
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
