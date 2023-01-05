package main

import "fmt"

func main() {
	// obj := Constructor()
	// res := obj.Get(1)
	// fmt.Println(res)

	obj := &MyLinkedList{Head: nil, Tail: nil}

	obj.AddAtHead(88)
	obj.PrintList()

	obj.AddAtTail(99)
	obj.PrintList()

	obj.AddAtIndex(5, 22)
	obj.PrintList()

	obj.DeleteAtIndex(2)
	obj.PrintList()
}

type MyLinkedList struct {
	Head, Tail *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func Constructor() MyLinkedList {
	node := &ListNode{Val: 999, Next: nil, Prev: nil}
	ll := MyLinkedList{Head: node, Tail: node}
	return ll
}

func (this *MyLinkedList) Get(index int) int {
	curr := this.Head
	loc := 0
	for curr.Next != nil {
		if index == loc {
			return curr.Val
		}
		curr = curr.Next
		loc++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Head == nil {
		node := &ListNode{Val: val, Next: nil, Prev: nil}
		this.Head = node
		this.Tail = node
		return
	}
	newHead := &ListNode{Val: val, Next: nil, Prev: nil}
	newHead.Next = this.Head
	this.Head.Prev = newHead
	this.Head = newHead
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.Head == nil {
		node := &ListNode{Val: val, Next: nil, Prev: nil}
		this.Head = node
		this.Tail = node
		return
	}
	newTail := &ListNode{Val: val, Next: nil, Prev: nil}
	newTail.Prev = this.Tail
	this.Tail.Next = newTail
	this.Tail = newTail
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.Head == nil && index != 0 {
		return
	}
	curr := this.Head
	loc := 0

	newNode := &ListNode{Val: val, Next: nil, Prev: nil}
	for curr.Next != nil {
		if index == loc+1 {
			temp := curr.Next

			curr.Next = newNode
			newNode.Next = temp

			temp.Prev = newNode
			newNode.Prev = curr
		}
		curr = curr.Next
		loc++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Head == nil {
		return
	}
	curr := this.Head
	loc := 0

	for curr.Next != nil {
		if index == loc {
			curr.Prev.Next = curr.Next
			curr.Next.Prev = curr.Prev
		}
		curr = curr.Next
		loc++
	}
}

func (this *MyLinkedList) PrintList() {
	curr := this.Head
	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
	fmt.Println()
}
