package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1}
	// k := 1
	nums := []int{-7, -8, 7, 5, 7, 1, 6, 0}
	k := 4
	// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// k := 2
	fmt.Println(maxSlidingWindow2(nums, k))
}

// 滑动窗口最大值，困难，自己能写出较为简单的，过不了时间限制。模仿可以写出高阶的，也过不了时间现在
// https://www.programmercarl.com/0239.%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%9C%80%E5%A4%A7%E5%80%BC.html
// 技巧：在滑动窗口中使用单调队列，保证队列中值从大到小，需要新做一个struct，这道题可以重新做一次，对新造struct很有帮助
// 这里的单调队列，只用于本题，如果题目中需要使用队列，可以根据自己的需要来做一个struct

// 高阶解法：自己设置一个单调队列，保证队列从大往小排列
func maxSlidingWindow2(nums []int, k int) []int {
	res := []int{}
	q := Constructor()
	for i := 0; i < len(nums); i++ {
		if i < k {
			// 在k之前，将数都加进去
			q.Push(nums[i])
			if i == k-1 {
				res = append(res, q.Out())
			}
			continue
		} else {
			// 将滑动窗口以前的值从queue中去掉
			q.Pop(nums[i-k])
			// 将滑动窗口新的值加进queue
			q.Push(nums[i])
			res = append(res, q.Out())
		}

	}
	return res
}

type Queue struct {
	queue []int
}

func Constructor() *Queue {
	return &Queue{
		queue: []int{},
	}
}

func (q *Queue) Out() int {
	return q.queue[0]
}

func (q *Queue) In() int {
	return q.queue[len(q.queue)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) Push(val int) {
	for !q.IsEmpty() && val > q.In() {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, val)
}

func (q *Queue) Pop(val int) {
	if !q.IsEmpty() && val == q.Out() {
		q.queue = q.queue[1:len(q.queue)]
	}
}

// 自己解法，用双指针，大部分可以过，对于超级大数字会超出时间限制
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= k {
		return []int{max(nums)}
	}

	inArr := false
	res := []int{}
	for curr := 0; curr <= len(nums)-k; curr++ {
		if curr == 0 {
			n := max(nums[curr : curr+k])
			res = append(res, n)
			inArr = n != nums[curr]
		}

		if curr != 0 {
			if nums[curr+k-1]-res[len(res)-1] > 0 {
				res = append(res, nums[curr+k-1])
				inArr = true
			} else if nums[curr+k-1]-res[len(res)-1] <= 0 && !inArr {
				n := max(nums[curr : curr+k])
				res = append(res, n)
				inArr = n != nums[curr]
			} else if nums[curr+k-1]-res[len(res)-1] <= 0 && inArr {
				res = append(res, res[len(res)-1])
				inArr = res[len(res)-1] != nums[curr]
			}
		}
	}
	return res
}

func max(arr []int) int {
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)
	return temp[len(temp)-1]
}
