package main

import "fmt"

func main23() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum2(nums, target))
}

// 两数之和，简单，自己做出来了
// https://www.programmercarl.com/0001.%E4%B8%A4%E6%95%B0%E4%B9%8B%E5%92%8C.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
// 重点：要想到用哈希表，map用来储存什么，key和value分别是什么

// 参考的，可以for loop一次就解决储存+检查
func twoSum2(nums []int, target int) []int {
	record := map[int]int{}
	for index, num := range nums {
		preIndex, ok := record[target-num]
		if ok {
			return []int{index, preIndex}
		} else {
			record[num] = index
		}
		fmt.Println(record)
	}
	return nil
}

// 自己做的，比较笨拙，需要存一次map，再检查一次map
func twoSum(nums []int, target int) []int {
	record := map[int][]int{}
	for i, num := range nums {
		record[num] = append(record[num], i)
	}

	for i := range record {
		indexA := record[i]
		want := target - i
		record[i] = record[i][1:]
		indexB, ok := record[want]
		if ok && len(indexB) != 0 {
			return []int{indexA[0], indexB[0]}
		}
	}
	return nil
}
