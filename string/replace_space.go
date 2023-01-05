package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "i love tanpu"
	fmt.Println(replaceSpace2(s))
}

// 替换字符串，简单
// https://www.programmercarl.com/%E5%89%91%E6%8C%87Offer05.%E6%9B%BF%E6%8D%A2%E7%A9%BA%E6%A0%BC.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 遍历添加
func replaceSpace2(s string) string {
	res := []rune{}
	for _, n := range s {
		if n == ' ' {
			res = append(res, []rune{'%', '2', '0'}...)
		} else {
			res = append(res, n)
		}
	}
	return string(res)
}

// 使用现成的库
func replaceSpace(s string) string { // 替换为%20
	res := strings.Replace(s, " ", "%20", -1)
	return res
}
