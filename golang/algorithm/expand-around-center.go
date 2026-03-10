package main

import (
	"fmt"
)

func main() {

	s0 := "babad"
	s1 := "abanoonttacxyzzzyx"
	s2 := "bb"

	fmt.Println(do(s0))
	fmt.Println(do(s1))
	fmt.Println(do(s2))

}

/*
aba → a|b|a
abbd → a|b|b|d

把字符间隙视作一个单位，那么aba奇数长度拥有偶数个间隙，abba偶数长度拥有奇数个间隙
n长度的字符串总共单位即2*n-1

中心扩展算法就是以此遍历字符+字符空隙。在每一次遍历过程中，通过左右双指针同时向两侧扩展，并校验是否符合回文规则
当不符合时，左右指针之间的距离，即是某一字符作为回文中心时，回文子串的最大长度-1
*/
func do(s string) string {
	n := len(s)

	if n < 2 {
		return s
	}

	maxLen := 1
	start := 0
	// 把字符间隙视作一个单位，那么aba奇数长度拥有偶数个间隙，abba偶数长度拥有奇数个间隙
	// n长度的字符串总共单位即2*n-1
	for i := 0; i < 2*n-1; i++ {
		// 在某次遍历时刻，取i的中间值（左右指针），分为两种情况i为奇数或偶数，通过取余的方式巧妙统一case
		left := i / 2
		right := left + i%2

		// inner for中校验回文规则，并由中心向两侧扩展
		for left >= 0 && right < n && s[left] == s[right] {
			// 这里实际上又是摒弃了空隙的单位，因此最终right - left -1的结果不用除以2
			left--
			right++
		}
		temp := right - left - 1
		if temp > maxLen {
			maxLen = temp
			// 跳出inner for时，left指针已经往左侧偏移了，因此此处需要+1
			start = left + 1
		}
	}
	return s[start : start+maxLen]
}
