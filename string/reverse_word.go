package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " asdasd df f  "
	fmt.Println("|", reverseWords2(s), "|")
}

// 将字符串中的单词翻转，中等，自己做出了暴力破解
// 一共3种解法：用库，暴力解法，原地修改考验功底解法

// 技巧：整体翻转+局部翻转。先去掉空格=>然后每个单词翻转=>然后整个列表翻转
// 自己写的，内存O（1），在列表上直接修改，涉及到多处双指针，可以好好做一下
func reverseWords3(s string) string {

	if len(s) == 0 {
		return ""
	}

	//去除首部空格
	i := 0
	for s[i] == ' ' {
		i++
	}
	s = s[i:]

	// 去除中间多余空格
	slow, fast, end := 0, 0, 0
	setEnd := false
	for ; fast < len(s); fast++ {
		if s[fast] == ' ' {
			if !setEnd {
				end = fast - 1
				setEnd = true
			}
		}

		if s[fast] != ' ' {
			if setEnd {
				s = s[:end+2] + s[fast:]
				slow = end + 2
				fast = slow
				setEnd = false
			}
		}
	}

	// 去除尾部多余空格
	j := len(s) - 1
	for s[j] == ' ' {
		j--
	}
	s = s[:j+1]

	// 翻转每个单词
	ss := []byte(s)

	slow, fast = 0, 0
	for fast < len(s) {
		if fast == len(s)-1 {
			reverseEachWord(ss[slow:len(s)])
			break
		}
		if s[fast] != ' ' {
			fast++
			continue
		}

		reverseEachWord(ss[slow:fast])
		slow, fast = fast+1, fast+1
	}

	// // 翻转这个列表
	left, right := 0, len(ss)-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}

	return string(ss)
}

func reverseEachWord(s []byte) {
	slow := 0
	fast := len(s) - 1
	for slow < fast {
		if s[fast] != ' ' {
			s[slow], s[fast] = s[fast], s[slow]
		}
		slow++
		fast--
	}
}

// 解法2：用库
func reverseWords2(s string) string {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	arr := strings.Split(s, " ")

	res := ""
	for i := len(arr) - 1; i >= 0; i-- {
		if i > 0 {
			res += arr[i] + " "
		} else if i == 0 {
			res += arr[i]
		}
	}
	return res
}

// 解法1：不用库的暴力破解
func reverseWords(s string) string {
	arr := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if i == len(s)-1 {
				s = s[:i]
				continue
			}
			str := s[i+1:]
			arr = append(arr, str)
			s = s[:i]
		}
		if i == 0 && len(s) != 0 {
			arr = append(arr, s)
		}
	}

	res := ""
	for i := 0; i < len(arr); i++ {
		if i < len(arr)-1 {
			res += arr[i] + " "
		} else if i == len(arr)-1 {
			res += arr[i]
		}
	}
	return res
}
