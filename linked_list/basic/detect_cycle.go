package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var cycle *ListNode
	cycle = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}
	curr := cycle
	for curr != nil {
		if curr.Next == nil {
			curr.Next = cycle
			break
		}
		curr = curr.Next
	}

	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: cycle}}
	fmt.Println(detectCycle2(head1))
}

// hasCycle判断是否有环，简单，自己做出来的，用快慢指针
// detectCycle找到环的位置，中等

// 题1：判断是否有环
// https://leetcode.cn/problems/linked-list-cycle/
// 技巧：使用快慢指针，一个每次2步，一个每次1步，如果存在环总会相遇
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 题2：找环的位置
// https://leetcode.cn/problems/linked-list-cycle-ii/comments/
// 解法1：先确定是否有环，然后参照get_intersection_node，不断循环两个列表，直至找到交点，
func detectCycle(head *ListNode) *ListNode {
	has, meet := hasCycleHelper(head)
	if !has {
		fmt.Println("come")
		return nil
	}

	curr1 := meet
	curr2 := head
	if meet == head {
		return meet
	}
	for curr1 != curr2 {
		if curr1.Next != meet {
			curr1 = curr1.Next
		} else {
			curr1 = head
		}

		if curr2.Next != meet {
			curr2 = curr2.Next
		} else {
			curr2 = meet
		}

		if curr1 == curr2 {
			return curr1
		}
	}

	return nil
}

// 解法2：先确定是否有环，然后根据公式得出，meet -> 环起点 == head -> 环起点
func detectCycle2(head *ListNode) *ListNode {
	has, meet := hasCycleHelper(head)
	if !has {
		return nil
	}
	if meet == head {
		return meet
	}
	node := head
	for node != nil {
		node = node.Next
		meet = meet.Next
		if node == meet {
			return node
		}
	}
	return nil
}

func hasCycleHelper(head *ListNode) (bool, *ListNode) {
	if head == nil || head.Next == nil {
		return false, nil
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, fast
		}
	}
	return false, nil
}
