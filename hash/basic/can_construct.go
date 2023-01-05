package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println(canConstruct2(ransomNote, magazine))
}

// 赎金信，判断一个字符串是否完全包含了另一个字符串，简单，自己做出来
// https://www.programmercarl.com/0383.%E8%B5%8E%E9%87%91%E4%BF%A1.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 使用数组，参考代码随想录
func canConstruct2(ransomNote string, magazine string) bool {
	record := make([]int, 26) // 也可以record := [26]int{}
	for _, l := range magazine {
		record[l-'a']++
	}
	fmt.Println(record)
	for _, l := range ransomNote {
		if record[l-'a'] >= 1 {
			record[l-'a']--
		} else if record[l-'a'] == 0 {
			return false
		}
	}
	return true
}

// 使用map，自己做的版本
func canConstruct(ransomNote string, magazine string) bool {
	record := map[rune]int{}
	for _, b := range magazine {
		count := record[b]
		record[b] = count + 1
	}

	for _, b := range ransomNote {
		count, ok := record[b]
		if ok && count > 0 {
			record[b]--
		} else {
			return false
		}
	}
	return true
}
