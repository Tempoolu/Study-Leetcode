package main

import (
	"fmt"
)

func main() {
	// arr := []int{-1, 0, 1, 2, -1, -4, 5, -5}
	arr := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	fmt.Println(threeSum(arr))
}

// 三数之和，中等，自己做的暴力解法
// https://www.programmercarl.com/0015.%E4%B8%89%E6%95%B0%E4%B9%8B%E5%92%8C.html#%E5%8F%8C%E6%8C%87%E9%92%88

// 只写出了暴力解法，做不出双指针写法，之后再来

// 高阶写法，双指针
// func threeSum2(nums []int) [][]int {
// 	sort.Ints(nums)
// 	left := 0
// 	right := len(nums) - 1
// 	i := left + 1

// 	res := [][]int{}
// 	for right-left >= 2 {
// 		fmt.Println(left, "  ", right, "  ", i)
// 		if nums[left]+nums[right]+nums[i] == 0 {
// 			res = append(res, []int{nums[left], nums[right], nums[i]})
// 			right--
// 			if i == right {
// 				left++
// 				i = left + 1
// 			}
// 		} else if nums[left]+nums[right]+nums[i] < 0 {
// 			i++
// 			if i == right {
// 				left++
// 				i = left + 1
// 			}
// 		} else if nums[left]+nums[right]+nums[i] > 0 {
// 			right--
// 		}
// 	}
// 	return res
// }

// 暴力解法，自己写的
func threeSum(nums []int) [][]int {

	fmt.Println(nums)
	res := [][]int{}

	// 应该学四数相加，用2+1记录，可以减少一层for loop操作
	for i := 0; i < len(nums)-2; i++ {
		for j := i; j < len(nums)-1; j++ {
			for k := j; k < len(nums); k++ {
				fmt.Println(i, " ", j, " ", k)
				if (i != j && j != k && i != k) && (nums[i]+nums[j]+nums[k] == 0) {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	// 用于去重，笨拙
	for i := 0; i < len(res)-1; i++ {
		for j := i; j < len(res); j++ {
			if sameArr(res[i], res[j]) {
				res = append(res[:i], res[i+1:]...)
			}
		}
	}
	return res
}

func sameArr(arrA, arrB []int) bool {
	record := map[int]int{}
	for _, n := range arrA {
		count := record[n]
		record[n] = count + 1
	}
	for _, n := range arrB {
		count, ok := record[n]
		if !ok && count == 0 {
			return false
		} else if ok && count >= 1 {
			record[n]--
		}
	}
	return true
}
