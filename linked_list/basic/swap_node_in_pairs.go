package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head1 := &ListNode{Val: 1, Next: nil}
	// head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	// head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	head3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	fmt.Println(swapPairs(head3))
	// printList(head1)
}

// 交换相邻节点，中等，自己做出来
// 技巧：使用虚拟头，可以不用管太多东西
// 顺序：先确定主要逻辑代码行，然后确定逻辑顺序，然后找边界（curr=nil,curr.Next=nil等，在for里面也可以设定边界)，然后找return什么（是否需要使用虚拟头）

func swapPairs(head *ListNode) *ListNode {
	curr := head
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: head.Next}
	prev := &ListNode{Val: 0, Next: nil}
	var next, nextNext *ListNode

	for curr != nil {
		if curr.Next == nil {
			break
		}
		if curr.Next != nil && curr.Next.Next == nil {
			next = curr.Next
			prev.Next = next
			next.Next = curr
			curr.Next = nil
			break
		}

		nextNext = curr.Next.Next
		next = curr.Next
		prev.Next = next
		next.Next = curr
		curr.Next = nextNext

		prev = curr
		curr = nextNext

	}
	// printList(dummy.Next)
	return dummy.Next
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
