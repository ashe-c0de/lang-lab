package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
*/
func main() {
	var str = "abcabcbb"
	var str2 = "你好世界红豆泥世界"
	var str3 = ""
	fmt.Println(test(str))
	fmt.Println(test(str2))
	fmt.Println(test(str3))
}

func test(str string) int {
	runes := []rune(str)
	m := make(map[rune]int)
	left := 0
	result := 0
	// 滑动窗口算法的核心就是，右指针给窗口扩容，直至抵达扩容限制条件或抵达边界；左指针则是给窗口缩容，以释放限制条件的约束，保证窗口继续向边界移动。
	for right := 0; right < len(runes); right++ {
		if v, ok := m[runes[right]]; ok && v >= left {
			left = v + 1
		}
		m[runes[right]] = right
		result = max(result, right-left+1)
	}

	return result
}
