package main

import (
	"fmt"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
*/
func main() {
	var str0 = "1221"
	//var str1 = "abcabcbb"
	//var str2 = "你好世界红豆泥世界"
	//var str3 = ""
	fmt.Println(Do(str0))
	//fmt.Println(Do(str1))
	//fmt.Println(Do(str2))
	//fmt.Println(Do(str3))
}

// 滑动窗口的核心就是，右指针给窗口扩容，直至抵达扩容限制条件或抵达边界；左指针则是给窗口缩容，以释放限制条件的约束，保证窗口继续向边界移动。
func Do(str string) int {
	var res, left int
	slice := []rune(str)
	m := make(map[rune]int) // key: 字符串字符, value: 字符所在字符串的index
	for right := 0; right < len(slice); right++ {
		// && val > left 是为了保持left是递增的，否则遇到重复字符left会变小进而res变大
		if val, ok := m[slice[right]]; ok && val > left {
			// 重复字符，左窗口缩容
			left = val + 1
		}
		m[slice[right]] = right
		res = max(res, right-left+1)
	}
	return res
}
