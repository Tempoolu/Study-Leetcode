package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

// 前N个高频元素，做不出来，主要考察各种排序

// 小顶堆做法，很难。使用堆排序，不会
func topKFrequent2(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// 构建小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 用包写，我写不出来，主要是sort.Slice写不出来
func topKFrequent(nums []int, k int) []int {
	record := map[int]int{}
	for _, n := range nums {
		amount := record[n]
		record[n] = amount + 1
	}
	fmt.Println(record)

	res := make([]int, k)
	for key, _ := range record {
		res = append(res, key)
	}

	// 这是重点
	// Slice sorts the slice x given the provided less function.
	sort.Slice(res, func(i, j int) bool {
		return record[res[i]] > record[res[j]]
	})
	return res[:k]
}
