package main

import "fmt"

func main() {
	fmt.Println(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

// 从array中移除特定元素，简单，自己做出来的
// https://leetcode.cn/problems/remove-element/submissions/
// 还有解法2：双指针法

// 方法1：暴力破解法
func removeElement(nums []int, val int) int {
	count := 0
	// 第一次for loop
	for i := 0; i < len(nums); i++ {
		count++
		if nums[i] == val && i < len(nums) {
			count--
			// 事实上，slice是array，append事实上是将后面的元素往前移动了，因此也算是O(N^2)
			// 第二次for loop
			nums = append(nums[:i], nums[i+1:]...)
			i--

		}
	}
	fmt.Println(nums)
	return count
}

// 方法2：快慢指针法
// https://www.programmercarl.com/0027.%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
func removeElement2(nums []int, val int) int {
	slow := 0

	// 快指针用于移动元素，慢指针用于更新数组
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[0:slow]
	fmt.Println(nums)
	return slow
}
