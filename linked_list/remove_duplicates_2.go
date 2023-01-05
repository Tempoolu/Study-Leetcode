package main

// 没做出来，甚至没理解到，中等难度
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

func deleteDuplicates2(head *ListNode) *ListNode {

	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
