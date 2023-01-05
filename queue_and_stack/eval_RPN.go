package main

import (
	"fmt"
	"strconv"
)

func main() {
	// s := []string{"4", "13", "5", "/", "+"}
	s := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(s))
}

// 逆波兰表达式求值，中等，自己做出来
// https://www.programmercarl.com/0150.%E9%80%86%E6%B3%A2%E5%85%B0%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%B1%82%E5%80%BC.html#%E9%A2%98%E5%A4%96%E8%AF%9D
// 主要是要想到用栈来做，这个也是类似于字符串消除，因此也应该首先想到栈

func evalRPN(tokens []string) int {
	stack := []string{}
	for i := 0; i < len(tokens); i++ {
		if isOps(tokens[i]) {
			fmt.Println(tokens[i])
			res := cal(strToNum(stack[len(stack)-2]), strToNum(stack[len(stack)-1]), tokens[i])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(res))
			fmt.Println(res)
		} else {
			stack = append(stack, tokens[i])
		}
	}
	return strToNum(stack[0])
}

func strToNum(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func isOps(str string) bool {
	ops := []string{"+", "-", "*", "/"}
	for _, s := range ops {
		if str == s {
			return true
		}
	}
	return false
}

func cal(num1, num2 int, ops string) int {
	switch ops {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0
}
