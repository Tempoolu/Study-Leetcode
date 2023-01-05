package main

import "fmt"

func main() {
	s := "abcdefghijklmnopq"
	fmt.Println(reverseStr2(s, 3))
}

// 翻转字符串2，简单，自己做出来但是花了很长时间
// https://www.programmercarl.com/0541.%E5%8F%8D%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
// 需要熟练掌握golang的string特性，以及map和slice对函数都是指针传递

func reverseStr2(s string, k int) string {
	ss := []byte(s) // 可以这样将string转为byte
	for i := 0; i < len(s); i += 2 * k {
		if len(s)-i > k {
			reverse(ss[i : i+k]) // 指针传递slice
		} else {
			reverse(ss[i:len(s)])
		}
	}
	return string(ss)
}

func reverse(s []byte) {
	left := 0
	right := len(s) - 1
	for right > left {
		s[left], s[right] = s[right], s[left]
		right--
		left++
	}
}

// 自己的罗里吧嗦版
func reverseStr(s string, k int) string {
	res := make([]byte, len(s), len(s))
	i := 0
	for i < len(s) {
		if len(s)-i < k && (i/k)%2 == 0 {

			left := i
			right := len(s) - 1
			for right >= i {
				res[left] = s[right]
				left++
				right--
			}
			break
		}
		if (i/k)%2 == 0 && i%k == 0 {
			left := i
			right := i + k - 1
			for right >= i {
				res[left] = s[right]
				left++
				right--
			}
			i += k
		} else {
			res[i] = s[i]
			i++
		}
	}
	return string(res)
}
