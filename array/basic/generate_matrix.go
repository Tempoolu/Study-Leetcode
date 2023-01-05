package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	// golang的二维数组需要这样建立
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// 以上

	curr := 1
	loc := 0
	end := n * n
	for curr <= end {
		curr = drawOutside(res, curr, n, loc)
		n -= 2
		loc += 1
	}
	return res
}

// curr：现在的数组，n：边长，loc：数组的位置
func drawOutside(res [][]int, curr, n, loc int) int {
	hor := loc // 二维数组的水平方向
	ver := loc // 二维数组的垂直方向

	if n == 1 {
		res[loc][loc] = curr
		curr++
	}

	for temp := 1; temp <= 4*(n-1); temp++ {
		if temp <= n-1 {
			res[ver][hor] = curr
			hor++
		} else if temp <= 2*n-2 {
			res[ver][hor] = curr
			ver++
		} else if temp <= 3*n-3 {
			res[ver][hor] = curr
			hor--
		} else if temp <= 4*n-4 {
			res[ver][hor] = curr
			ver--
		}
		curr++
	}

	return curr
}
