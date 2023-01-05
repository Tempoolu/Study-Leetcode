package main

import (
	"fmt"
)

func main() {

	q := Constructor1()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

type MyStack struct {
	in  []int
	out []int
}

func Constructor1() MyStack {
	return MyStack{
		in:  []int{},
		out: []int{},
	}
}

func (this *MyStack) Move() {
	if len(this.in) == 0 {
		return
	}
	for i := 0; i < len(this.in); i++ {
		this.out = append(this.out, this.in[i])
	}
	this.in = []int{}
}

func (this *MyStack) Push(x int) {
	this.in = append(this.in, x)
	this.Move()
}

func (this *MyStack) Pop() int {
	if len(this.out) == 0 {
		return -1
	}
	last := this.out[len(this.out)-1]     // return最后一个out
	this.out = this.out[:len(this.out)-1] // 将out的最后一个删除
	return last
}

func (this *MyStack) Top() int {
	last := this.Pop()
	this.Push(last)
	return last

}

func (this *MyStack) Empty() bool {
	if len(this.in) == 0 && len(this.out) == 0 {
		return true
	}
	return false
}
