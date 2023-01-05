package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// intersectionNode := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}

	// list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: intersectionNode}}
	// list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: intersectionNode}}}

	// list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: intersectionNode}}
	// list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: intersectionNode}}

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	fmt.Println(getIntersectionNode2(list1, list2))
}

// 获取相交节点，中等，自己做出来
// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/submissions/
// 两种解题思路

// 思路1：分别计算长度，然后用快慢指针相遇
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == headB {
		return headA
	}
	currA, currB := headA, headB
	var lenA, lenB int

	// 先计算两个列表的长度
	for currA != nil {
		currA = currA.Next
		lenA++
	}
	for currB != nil {
		currB = currB.Next
		lenB++
	}

	// 快慢指针分别遍历
	count := 1
	if lenA-lenB < 0 {
		// 交换A,B是为了让下面的代码不用写两次
		headA, headB, lenA, lenB = swap(headA, headB, lenA, lenB)
	}
	diff := lenA - lenB

	currA = headA
	currB = headB
	for currA != nil {
		if count > diff {
			currB = currB.Next

		}
		count++
		currA = currA.Next
		if currA == currB {
			return currA
		}
		fmt.Println(currA, "  ", currB)
	}
	return nil
}

func swap(headA, headB *ListNode, lenA, lenB int) (*ListNode, *ListNode, int, int) {
	tempHead := headB
	headB = headA
	headA = tempHead

	tempLen := lenB
	lenB = lenA
	lenA = tempLen

	return headA, headB, lenA, lenB
}

// 走完A就走B，走完B就走A
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	currA, currB := headA, headB
	// 如果没有相交，在currA==currB=nil时会推出for loop
	for currA != currB {
		if currA != nil {
			currA = currA.Next
		} else if currA == nil {
			currA = headB
		}

		if currB != nil {
			currB = currB.Next
		} else if currB == nil {
			currB = headA
		}
	}
	return currA
}
