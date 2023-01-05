package main

import (
	"fmt"
)

func main1() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

// 有序数组的平方，简单，使用左右双指针，以及创建新arr一个指针，知道概念自己可以写出来
// https://leetcode.cn/problems/squares-of-a-sorted-array/comments/
// https://www.programmercarl.com/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.html#%E6%9A%B4%E5%8A%9B%E6%8E%92%E5%BA%8F

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums), len(nums)) // 这里的res需要指定长度（所有数设为0），以及容量
	for i := len(nums) - 1; i >= 0; i-- {
		if pow2(nums[right]) > pow2(nums[left]) {
			res[i] = nums[right]
			right--
		} else if pow2(nums[right]) < pow2(nums[left]) {
			res[i] = pow2(nums[left])
			left++
		} else {
			res[i] = pow2(nums[left])
			break
		}
	}
	return res
}

func pow2(n int) int {
	return n * n
}
