package main

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func reverseList(head *ListNode2) *ListNode2 {
	if head == nil || head.Next == nil {
		return head
	}

	nh := reverseList(head.Next)
	head.Next.Next = head // head=4, 4.next = 5 意味着 5.Next = 4
	head.Next = nil       // head=4，指向空
	return nh
}
