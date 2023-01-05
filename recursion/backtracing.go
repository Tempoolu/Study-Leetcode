package main

import (
	"fmt"
	"sort"
)

// 全排列，回溯，和树有关，不会做
// https://bbs.huaweicloud.com/blogs/349853 //回溯的各种案例
// https://leetcode.cn/problems/permutations/solution/pei-yang-chou-xiang-neng-li-de-yi-dao-ti-1731/

func main1() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	sort.Ints(nums)
	// 回溯
	res := [][]int{}

	var backTrack func(path []int)
	backTrack = func(path []int) {

		if len(path) == len(nums) {
			// 把path添加到reszhong
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 如果当前路径包含相同字母，不添加
			// flag := false
			// for _, v := range path {
			// 	if v == nums[i] {
			// 		flag = true
			// 	}
			// }
			// if flag {
			// 	continue
			// }
			if isExist(path, nums[i]) {
				continue
			}

			path = append(path, nums[i])
			backTrack(path)
			path = path[:len(path)-1]
		}

	}
	backTrack([]int{})
	return res
}

// func permute(nums []int) [][]int {
// 	res := [][]int{}
// 	var backTrack func(path []int)
// 	backTrack = func(path []int) {

// 		if len(path) == len(nums) {
// 			temp := []int{}
// 			copy(temp, path)
// 			fmt.Println("temp is ", temp)
// 			res = append(res, temp)
// 			return
// 		}

// 		for i := 0; i < len(nums); i++ {
// 			if isExist(path, nums[i]) {
// 				continue
// 			}
// 			path = append(path, nums[i])
// 			fmt.Println("path is ", path)
// 			backTrack(path)
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	backTrack([]int{})
// 	fmt.Println(res)
// 	return res
// }

func isExist(arr []int, target int) bool {
	for _, i := range arr {
		if i == target {
			return true
		}
	}
	return false
}
