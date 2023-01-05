package main

import (
	"fmt"
)

func main() {

	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())
	fmt.Println(q.Empty())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

type MyQueue struct {
	in  []int
	out []int
}

// 用两个栈来实现队列，简单，自己做出来
// https://www.programmercarl.com/0232.%E7%94%A8%E6%A0%88%E5%AE%9E%E7%8E%B0%E9%98%9F%E5%88%97.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
// 重点是：pop中，如果out为空，需将in移到out中再pop
// 此外，peek要复用pop的代码是更好的习惯

func Constructor() MyQueue {
	return MyQueue{
		in:  []int{},
		out: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		}
		this.out = make([]int, len(this.in))
		i, o := len(this.in)-1, 0
		for i >= 0 {
			this.out[o] = this.in[i]
			i--
			o++
		}
		this.in = []int{}
	}

	last := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return last

}

// Peek复用版本
func (this *MyQueue) Peek() int {
	last := this.Pop()
	this.out = append(this.out, last)
	return last
}

func (this *MyQueue) Empty() bool {
	if len(this.in) == 0 && len(this.out) == 0 {
		return true
	}
	return false
}

// 这个是自己写的peek，重新写了一次，应该复用pop
func (this *MyQueue) Peek2() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		} else {
			return this.in[0]
		}
	}
	return this.out[len(this.out)-1]
}
