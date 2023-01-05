package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 涉及到反转链表和快慢指针，可以重新做
// 是我自己做出来的！
// https://leetcode.cn/problems/palindrome-linked-list/

// 技巧：使用快慢指针找到中点，一边找中点一边反转链表
func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	var pre *ListNode
	for fast.Next != nil {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next

		if fast == nil {
			break
		}
	}

	var left, right *ListNode
	if fast == nil {
		left = pre
		right = slow
	} else {
		left = pre
		right = slow.Next
	}

	for left != nil {
		if left.Val == right.Val {
			left = left.Next
			right = right.Next
			continue
		} else {
			return false
		}
	}
	return true
}
