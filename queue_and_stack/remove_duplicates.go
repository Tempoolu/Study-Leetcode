package main

import "fmt"

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates(s))

}

// 消除相邻相同字母，简单，自己做出来
// https://www.programmercarl.com/1047.%E5%88%A0%E9%99%A4%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%B8%AD%E7%9A%84%E6%89%80%E6%9C%89%E7%9B%B8%E9%82%BB%E9%87%8D%E5%A4%8D%E9%A1%B9.html#%E6%AD%A3%E9%A2%98
// 如果是相邻消除，优先考虑栈。递归是用栈来实现的。
// 拓展：未来可以尝试用双指针完成

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else if len(stack) != 0 {
			if s[i] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	return string(stack)
}
