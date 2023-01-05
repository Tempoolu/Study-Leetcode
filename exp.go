package main

import "fmt"

// 需要学习堆排序，kmp，

func main1() {
	// 关于rune
	// a := rune('a')
	// fmt.Println(a) // 97

	// 关于循环map
	// b := map[int]bool{}
	// b[3] = true
	// for key, value := range b {
	// 	fmt.Println("key is ", key)
	// 	fmt.Println("value is ", value)
	// }

	// 关于列表指针
	// a := []int{1, 2, 3}
	// changeOneNum(a[0]) // 试图只修改一个num
	// fmt.Println(a)     // 依然是1，2，3

	// changeSlice(a) // 修改整个slice，会将指针传给函数
	// fmt.Println(a) // 改成了0,2,3

	// var changeInside func(int)
	// changeInside = func(i int) {
	// 	i = 4
	// }
	// changeInside(a[0])
	// fmt.Println(a) // 依然是0,2,3

	// 关于string指针
	// a := "abc"
	// changeLetterValue(a)
	// fmt.Println(a) // abc
	// changeLetterAddress(&a)
	// fmt.Println(a) // ooo

	// 关于string遍历rune和byte
	// a := "我爱谭噗"

	// for i := 0; i < len(a); i++ { // utf-8遍历，byte，输出的是一个字节
	// 	fmt.Println(string(a[i])) // 乱码
	// 	fmt.Println(a[i])         // 一堆数字
	// }

	// for _, n := range a { // unicode遍历，rune，输出的是四个字节
	// 	fmt.Println(string(n)) // 我 爱 谭 噗
	// 	fmt.Println(n)         // 25105 29233 35885 22103
	// }

	// fmt.Println(rune('a'), " ", rune('A')) // 97 65
	// fmt.Println(byte('a'), " ", byte('A')) // 97 65
	// // fmt.Println(byte('谭')) // 不可以会报错
	// fmt.Println(rune('谭')) // 35885

	// 关于位数
	// fmt.Println(123 / 10) // 12
	// fmt.Println(123 % 10) // 3

	// fmt.Println(9 / 10) // 0
	// fmt.Println(9 % 10) // 9

	// 关于字典
	// dict := map[int]int{}
	// count, ok := dict[1]
	// fmt.Println("count is ", count, ", ok is ", ok)

	// 关于string的指针
	// str := "abc"
	// fmt.Println(&str) // 0xc000088220
	// str = str[:1]
	// fmt.Println(&str) // 0xc000088220 属于原地操作

	// b := []byte(str)
	// fmt.Println(&b)    // &[97] 是链表
	// fmt.Println(&b[0]) // 0xc000014088

	// 关于copy
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(b)
}

func changeOneNum(num int) {
	num = 0
}
func changeSlice(arr []int) {
	arr[0] = 0
}

func changeLetterValue(s string) {
	s = "ooo"
}

func changeLetterAddress(s *string) {
	*s = "ooo"
}
