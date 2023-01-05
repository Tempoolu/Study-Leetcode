package main

import "fmt"

// 递归最重要的就是————————递归公式+边界条件

func main() {
	// fmt.Println(multiplier(5))
	// fmt.Println(fic(5))
	// fmt.Println(add([]int{1, 2, 3}))
	// fmt.Println(max([]int{1, 3, 2, 4, 1, 5, 7, 3}, 0))
	// fmt.Println(addDigit(9875))
	fmt.Println(powerOfTwo(1))
}

// 乘阶
func multiplier(n int) int {
	if n == 1 { // 边界条件
		return 1
	}

	return n * multiplier(n-1) // 递归公式
}

// Fibonacci数列
func fic(n int) int {
	if n == 0 || n == 1 { // 边界条件
		return 1
	}
	return fic(n-1) + fic(n-2) // 递归公式
}

// 给一个数列，用递归求和
func add(arr []int) int {
	if len(arr) == 1 { // 边界条件
		return arr[0]
	}
	return arr[len(arr)-1] + add(arr[0:len(arr)-1]) // 递归公式
}

// 给一个数列，用递归求最大值
func max(arr []int, big int) int {
	if len(arr) == 1 { // 边界条件
		if arr[0] > big {
			return arr[0]
		} else {
			return big
		}
	}

	if arr[0] > big { // 递归公式
		return max(arr[1:], arr[0])
	} else {
		return max(arr[1:], big)
	}
}

// 将一个数字每个位数一直相加，直到数字为0为止
// https://leetcode.cn/problems/add-digits/solution/by-jigexio-sd7m/
func addDigit(num int) int {
	if num%10 == num {
		return num
	}

	num = addDigit(num%10) + addDigit(num/10)

	if num%10 != num {
		num = addDigit(num)
	}

	return num
}

// 判断数字是否是2的幂
// https://leetcode.cn/problems/power-of-two/solution/di-gui-wan-cheng-by-wyppeng-02ju/
func powerOfTwo(n int) bool {

	if n < 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%2 == 0 {
		return powerOfTwo(n / 2)
	}
	return false
}
