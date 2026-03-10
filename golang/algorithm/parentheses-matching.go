package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
输入：s = "()"
输出：true

输入：s = "()[]{}"
输出：true

输入：s = "(]"
输出：false

输入：s = "([])"
输出：true

输入：s = "([)]"
输出：false
*/
func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
	s2 := "([])"
	fmt.Println(isValid(s2))
	s3 := "([)]"
	fmt.Println(isValid(s3))
	s4 := ")("
	fmt.Println(isValid(s4))
}

// “括号匹配算法” (Parentheses Matching Algorithm)
func isValid(str string) bool {
	strLen := len(str)
	if strLen%2 != 0 {
		return false
	}
	hmap := map[byte]byte{
		'(': ')', // 40 41
		'{': '}',
		'[': ']',
	}

	var stack []byte // 存放hmap的val,并在出栈时从末端截取（LIFO）
	for i := 0; i < len(str); i++ {
		if v, ok := hmap[str[i]]; ok {
			// 入栈
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != str[i] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
