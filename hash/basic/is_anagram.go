package main

import "fmt"

func main() {
	a := "anagram"
	b := "nagaram"
	fmt.Println(isAnagram3(a, b))
}

// 检查s和t每个字母出现的次数是否相同，简单，自己做出来
// https://leetcode.cn/problems/valid-anagram/

// 解法1：两个表分别记录map，再逐一对比map，O(3n)，对比map比较多余
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		count := mapS[s[i]]
		mapS[s[i]] = count + 1
	}

	mapT := make(map[byte]int, 26)
	for i := 0; i < len(t); i++ {
		count := mapT[t[i]]
		mapT[t[i]] = count + 1
	}

	for a := range mapS {
		if mapS[a] != mapT[a] {
			return false
		}
	}
	return true
}

// 解法2：for loop1循环出一个map，再for loop注意删掉map的元素
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		count := record[s[i]]
		record[s[i]] = count + 1
	}

	for i := 0; i < len(t); i++ {
		count, ok := record[t[i]]
		if count == 0 || !ok {
			return false
		}
		record[t[i]] = count - 1
	}

	return true
}

// 解法3：用列表记录
func isAnagram3(s string, t string) bool {
	record := [26]int{}
	for _, r := range s {
		record[r-rune('a')]++ //如果record[r] = record[97]，就超出range报错，因此减去rune('a')
	}

	for _, r := range t {
		record[r-rune('a')]--
	}
	return record == [26]int{}
}
