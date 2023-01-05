package main

import "fmt"

// 全排列，回溯，和树有关，不会做
// https://bbs.huaweicloud.com/blogs/349853 //回溯的各种案例
// https://leetcode.cn/problems/permutations/solution/pei-yang-chou-xiang-neng-li-de-yi-dao-ti-1731/

func main() {
	s := []byte{'t', 'a', 'n', 'p', 'u', 'q', 'i'}
	reverseString(s)
}

// 反转字符串，简单，自己做出来
// https://leetcode.cn/problems/reverse-string/submissions/
// 注意边界条件

func reverseString(s []byte) {
	fmt.Println(s)
	fmt.Println(len(s) / 2)
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}
