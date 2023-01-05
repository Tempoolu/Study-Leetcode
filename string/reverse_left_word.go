package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	fmt.Println(reverseLeftWords(s, 2))
}

// 左旋转字符串，简单，高阶要求在原地修改，自己写出来
// https://www.programmercarl.com/%E5%89%91%E6%8C%87Offer58-II.%E5%B7%A6%E6%97%8B%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
// 技巧：整体旋转和原地旋转

func reverseLeftWords(s string, n int) string {
	// 在golang中，如果要替换字符串中的byte，就一定要变成[]byte，需要多分配一个内存
	// 但是如果只是修建，s=s[:n]+s[n+2:]，就不需要重新分配内存
	ss := []byte(s)

	reverse(ss[:n])
	reverse(ss[n:])
	reverse(ss)
	return string(ss)
}

func reverse(ss []byte) {
	left, right := 0, len(ss)-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		right--
		left++
	}
}
