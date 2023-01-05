package main

import (
	"fmt"
)

func main2() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

// 解法一：暴力破解，O(n^2)
func minSubArrayLen(target int, nums []int) int {
	ans := 99999
	for i := range nums {
		add := 0
		for j := i; j < len(nums); j++ {
			add += nums[j]
			if add >= target {
				if j-i+1 < ans {
					ans = j - i + 1
				}
				break
			}
			if add > target {
				break
			}
		}
	}
	return ans
}

// 解法二：滑动窗口,O(n)
// https://www.programmercarl.com/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
func minSubArrayLen2(target int, nums []int) int {
	slow := 0 // 慢指针，指向窗口的开始位置
	sum := 0
	ans := len(nums) + 1

	for fast := 0; fast < len(nums); fast++ { // 快指针，指向窗口的结束位置
		sum += nums[fast]
		for sum >= target {
			// 如果窗口值太大，则将开始位置往后一个，for直至==target（记录）或<target（新一轮外部循环）
			sum -= nums[slow]
			subLen := fast - slow + 1
			if ans > subLen {
				ans = subLen
			}
			slow++
		}
	}

	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
