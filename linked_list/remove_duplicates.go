package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 自己做出来的，难度简单
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	cur := head
	pre := cur
	for cur.Next != nil {

		if cur.Val == cur.Next.Val {
			cur = cur.Next
		} else {
			cur = cur.Next
			pre.Next = cur
			pre = cur
		}
	}
	pre.Next = cur.Next

	return head
}
