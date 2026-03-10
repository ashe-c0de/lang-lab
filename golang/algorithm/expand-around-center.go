package main

import (
	"fmt"
)

/*
abad  → a|b|a|d
refer → r|e|f|e|r

中心扩展算法处理最长回文子串的核心理解在于：把字符和字符间的间隙（gap）抽象成一个单位，这样偶数字符串abad用奇数（3）个间隙，奇数字符串refer有偶数（4）个间隙。
n个长度的字符串，始终存在2n-1个单位

中心扩展算法就是以此遍历字符+字符空隙。在每一次遍历过程中，通过左右双指针同时向两侧扩展，并校验是否符合回文规则
当不符合时，左右指针之间的距离，即是某一字符作为回文中心时，回文子串的最大长度-1
*/
func longestPalindrome(s string) string {

	n := len(s)
	if n < 2 {
		return s
	}

	start := 0
	maxLen := 1

	// 共有 2n-1 个回文中心
	for center := 0; center < 2*n-1; center++ {

    // 回文中心的左右指针
		left := center / 2
		right := left + center%2

		// 向两边扩展
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}

		// 退出inner循环时已经越界，需要回退一步
		length := right - left - 1

		if length > maxLen {
			maxLen = length
			start = left + 1
		}
	}

	return s[start : start+maxLen]
}

func main() {

	s0 := "babad"
	s1 := "abanoonttacxyzzzyx"

	fmt.Println(longestPalindrome(s0))
	fmt.Println(longestPalindrome(s1))

}
