package main

import "fmt"

/*
判断一个字符串是否是回文字符串
*/
func main() {
	var s0 = "refer"
	var s1 = "refers"
	var s2 = "aa"
	var s3 = "你好世界捏界世好你"
	fmt.Println(do(s0))
	fmt.Println(do(s1))
	fmt.Println(do(s2))
	fmt.Println(do(s3))
}

// 双指针算法
func do(str string) bool {
	slice := []rune(str)
	right := len(slice) - 1
	var left int
	for left < right {
		if slice[left] != slice[right] {
			return false
		}
		left++
		right--
	}
	return true
}
