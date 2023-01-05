package main

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func main() {
	q := Constructor()

	q.Push(3)
	q.Push(5)
	q.Push(8)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	q.Peek()
	fmt.Println(q.Empty())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

func Constructor() MyQueue {
	return MyQueue{
		in:  []int{},
		out: []int{},
	}
}

func (this MyQueue) Pop() int {

	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		}
		this.out = make([]int, len(this.in))
		i, o := 0, len(this.in)-1
		for o > 0 {
			this.out[o] = this.in[i]
			o--
			i++
		}
	}

	if len(this.out) != 0 {
		n := this.out[len(this.out)-1]
		this.out = this.out[:len(this.out)-1]
		return n
	}
	return -1
}

func (this MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this MyQueue) Peek() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		}
	}
	return this.in[0]
}

func (this MyQueue) Empty() bool {
	if len(this.in) == 0 && len(this.out) == 0 {
		return true
	}

	return false
}
