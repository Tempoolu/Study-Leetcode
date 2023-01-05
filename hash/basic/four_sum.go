package main

import "fmt"

func main() {
	// nums1 := []int{1, 2}
	// nums2 := []int{-1, -2}

	// nums3 := []int{-1, 2}
	// nums4 := []int{0, 2}

	nums1 := []int{0}
	nums2 := []int{0}

	nums3 := []int{0}
	nums4 := []int{0}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}

// 四数相加，中等，看到思路自己写出来的
// https://www.programmercarl.com/0454.%E5%9B%9B%E6%95%B0%E7%9B%B8%E5%8A%A0II.html
// 技巧，将四个数两两分配，nums1和nums2一组，nums3和nums4一组，然后分别相加看是否为0

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	recordA, recordB := map[int]int{}, map[int]int{}
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			add := n1 + n2
			count := recordA[add]
			recordA[add] = count + 1
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			add := n3 + n4
			count := recordB[add]
			recordB[add] = count + 1
		}
	}

	ans := 0
	for sumA, countA := range recordA {
		countB, has := recordB[0-sumA]
		if has {
			ans += countA * countB
		}
	}

	return ans
}
