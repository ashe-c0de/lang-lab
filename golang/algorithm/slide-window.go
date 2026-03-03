/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
*/
func main() {
	var str = "abcabcbb"
	var str2 = "abcabcfbb"
	var str3 = ""
	fmt.Println(test(str))
	fmt.Println(test(str2))
	fmt.Println(test(str3))
}

func test(str string) int {
	m := make(map[byte]int)
	left := 0
	res := 0
	// 滑动窗口算法的核心就是，右指针给窗口扩容，直至抵达扩容限制条件或抵达边界；左指针则是给窗口缩容，以释放限制条件的约束，保证窗口继续向边界移动。
	for right := 0; right < len(str); right++ {
		if index, ok := m[str[right]]; ok && index >= left {
			left = index + 1
		}
		m[str[right]] = right
		res = max(res, right-left+1)
	}

	return res
}
