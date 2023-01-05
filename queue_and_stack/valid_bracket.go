package main

import (
	"fmt"
)

func main() {
	s := "(){(}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == m[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
