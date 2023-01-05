package main

import (
	"fmt"
)

func main5() {
	// nums1 := []int{-26, -11, 3, -8, -31, -29, -7, -10, -8, 9, -24, 9, -18, 2, -3, 5, 4, 1, -4, -23, -19, -14, -28, -30, -25, -15, 1, -27, -19, 9, 5, -15, -6, 8, 9, -23, -9, -14, 2, 3, -2, -4, 4, -15, 7, 9, -14, -8, -3, -1, -28, -2, -19, 2, -31, -8, -19, -9, -29, -15, -4, -2, -18, 10, 2, -8, -6, 4, 4, -25, 2, -27, -21, -12, -15, 2, 9, 4, -4, -18, -22, -18, -24, -20, 4, -6, -12, -4, 7, -23, -32, -4, -16, 5, 3, -21, -15, -22, 5, -25}
	// nums2 := []int{-29, -1, -12, -12, -4, 4, 4, -25, -21, -3, -25, -2, -16, -15, -22, -26, -31, -9, -19, 10, -1, -1, -6, -28, -32, -31, -10, -31, -8, 1, 5, -19, -29, -8, -25, -25, -14, -27, -1, -26, 6, -20, 0, -3, -25, -22, 3, -20, -25, -4, -6, -27, -8, -26, -4, -25, -31, -27, -32, -29, -14, 10, -25, -11, -32, -25, -13, -5, -27, -30, -5, -4, -12, 0, -4, 9, -27, -6, -28, -30, -19, -5, -19, -18, -3, -32, -17, -1, -10, 3, -17, -31, -26, -16, -10, 6, -19, -19, 3, -2}
	// nums3 := []int{-3, -17, -23, -28, -27, -28, -28, 5, -18, -32, -7, 0, -27, 5, -15, -30, 9, 6, 4, -11, 3, 5, 10, -2, -32, -22, -30, -9, -4, 10, 8, -9, 0, -26, -8, -7, -3, -21, -22, -28, -17, -14, -15, -9, 10, 4, -8, -32, 9, -7, -7, -18, -5, 2, -11, -27, -2, 9, 8, -26, -2, -8, -32, -11, -3, -16, -22, 8, -29, -3, -10, -30, 8, 5, -23, -2, -22, -15, -32, -2, -23, -3, 0, -16, -10, -29, -18, -24, -24, -16, -1, 3, -22, -23, 3, -20, 7, -16, -11, 6}
	// nums4 := []int{-4, -13, -10, -2, 5, -23, -10, 6, 7, -23, -12, -21, 4, -8, -9, -5, -18, -5, -22, -30, -23, 9, -27, -1, 5, -18, -23, -6, 1, -23, -20, -30, 5, -21, -15, -30, 10, -26, -7, -10, -21, -12, -4, 9, -8, 2, -27, -14, -11, 0, -1, -26, -27, 8, -14, 9, -19, -23, -10, -31, -16, -17, -22, -13, 9, -26, 5, -14, 4, -11, 4, -8, 4, -29, -32, -24, -22, -6, -3, 7, -17, -19, -32, 9, -31, -31, -30, -6, -10, -4, -7, -27, -25, -25, 7, -6, -17, 7, -30, 6}

	nums1 := []int{1, 2}
	nums2 := []int{-1, -2}
	nums3 := []int{0, 3}
	nums4 := []int{1, -1}

	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}

func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	return 0
}

// 暴力破解，但是再leetcode中超出时间限制
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			for _, n3 := range nums3 {
				for _, n4 := range nums4 {
					if n1+n2+n3+n4 == 0 {
						count++
					}
				}
			}
		}
	}
	return count
}
