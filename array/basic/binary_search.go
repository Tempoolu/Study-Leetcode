package main

import "fmt"

func main() {
	fmt.Println(binarySearch2([]int{-1, 0, 3, 5, 9, 12}, 9))
}

// 二分查找，非常基础
// https://leetcode.cn/problems/binary-search/

// 左闭右闭版本
func binarySearch1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		// [1,2,3,4]，第一次二分会从mid_index=1开始，[1,2]和[3,4]
		// [1,2,3,4,5]，第一次二分会从mid_index=2开始，[1,2,3]和[4,5]
		// 其实都没关系，只要nums[mid] != target，就不会出错

		mid := left + ((right - left) / 2) // 防止溢出，等同于(right+left)/2，事实上这道题两种都可以
		if nums[mid] > target {
			right = mid - 1 // 由于不可能是mid，因此-1
		}
		if nums[mid] < target {
			left = mid + 1 // 由于不可能是mid，因此+1
		}
		if nums[mid] == target {
			return mid
		}
	}
	return -1
}

// 左闭右开版本
func binarySearch2(nums []int, target int) int {
	left := 0
	right := len(nums) // 这个right不存在，因此left==right没有价值，因此下面left < right，而非left <= right
	for left < right {
		mid := left + ((right - left) / 2) // 防止溢出，等同于(right+left)/2，事实上这道题两种都可以
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] < target {
			left = mid
		}
		if nums[mid] == target {
			return mid
		}
	}
	return -1
}
