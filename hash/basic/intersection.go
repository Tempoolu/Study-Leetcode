package main

import (
	"fmt"
)

func main3() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{4, 5, 3}
	fmt.Println(intersection(nums1, nums2))
}

// 两个数组的交集，只输出一次，简单，自己做
// https://leetcode.cn/problems/intersection-of-two-arrays/submissions/
// 技巧，使用map[int]bool，还有如何用map实现set的效果

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}

	record := map[int]bool{}
	for _, num := range nums1 {
		record[num] = true
	}

	for _, num := range nums2 {
		has, ok := record[num]
		if has && ok {
			// golang没有se的概念，用这个方法来实现只存一次
			// record[num] = false

			// 也可以用删掉key来实现只存一次
			delete(record, num)
			res = append(res, num)
		}
	}

	return res
}
