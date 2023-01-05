package main

import "fmt"

func main6() {
	fmt.Println(isHappy(1))
}

// 判断快乐数，简单
// https://leetcode.cn/problems/happy-number/comments/
// 找到规律，规律是如果存在重复的结果，就会一直循环下去，无法获得1，因此用map当作set
// 其它：要熟悉位数相加的情况

func isHappy(n int) bool {
	var cal func(int) int
	cal = func(i int) int {
		ans := 0
		// 个位数/10=0
		for i > 0 {
			ans += (i % 10) * (i % 10)
			i = i / 10
		}
		return ans
	}
	record := map[int]bool{}
	for {
		// 顺序：先检查，再处理数字迭代
		_, ok := record[n]
		if ok {
			return false
		} else if n == 1 {
			return true
		}
		record[n] = true
		n = cal(n)
	}
}
