package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 5, 7, 3, 8, 0, 4}
	// fmt.Println(bubbleSort(arr))
	// fmt.Println(selectionSort(arr))
	// fmt.Println(inserctionSort(arr))
	// fmt.Println(mergeSort(arr))
	fmt.Println(quickSort(arr))
}

// 冒泡算法
// O(N^2)，适合特别小的数据量
// 比较相邻两个数值，将小的放在前面。不断重来，直到完全排好。每次内部循环，都会将最大数字放到最后，因此下一次循环无需比较最后一个数字。
func bubbleSort(arr []int) []int {
	var isSorted bool
	for i := len(arr) - 1; i >= 0; i-- {
		isSorted = false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				isSorted = true
			}
		}
		if isSorted == false {
			break
		}
	}
	return arr
}

// 选择排序
// 0(N^2)，适合小的数据量，是冒泡算法的优化
// 内循环找出最小的值，将第一个值和最小值交换。然后不断外循环，直到结束
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		small := i
		for j := i; j < len(arr); j++ {
			if arr[j] < small {
				small = j
			}
		}
		arr[i], arr[small] = arr[small], arr[i]
	}
	return arr
}

// 插入排序
// O(N^2)，适合小的数据量
// 分为左边已排序，右边未排序，将右边的第一个数字插入左边已排序好的，不断循环
func inserctionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

// 归并排序
// 在数据量较大的时候表现不错，但空间复杂度为O(n)，因此在数据量上千万的时候很占内存
// 使用迭代or递归，将数据不断拆分排序，8=>4=>2排序
func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return sortArrForMerge(left, right)
}

func sortArrForMerge(lArr []int, rArr []int) []int {
	lIndex, rIndex := 0, 0
	res := []int{}
	for lIndex < len(lArr) && rIndex < len(rArr) {
		if lArr[lIndex] < rArr[rIndex] {
			res = append(res, lArr[lIndex])
			lIndex++
		} else {
			res = append(res, rArr[rIndex])
			rIndex++
		}
	}
	if rIndex < len(rArr) {
		res = append(res, rArr[rIndex:]...)
	}
	if lIndex < len(lArr) {
		res = append(res, lArr[lIndex:]...)
	}
	return res
}

func quickSort(arr []int) []int {
	arr = []int{4, 3, 2, 5, 1}
	return sortArrForQuick(arr)
}

func sortArrForQuick(arr []int) []int {
	// pivot := arr[0]
	i, j := 0, len(arr)-1
	for i != j {
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		} else {
			j--
		}
	}
	return arr
}
