/*
Given a string s, find the longest substring which is a palindrome. If there are multiple answers, then find the first appearing substring.


Input: s = "forgeeksskeegfor"
Output: "geeksskeeg"
Explanation: The longest substring that reads the same forward and backward is "geeksskeeg". Other palindromes like "kssk" or "eeksskee" are shorter.

Input: s = "Geeks"
Output: "ee"
Explanation: The substring "ee" is the longest palindromic part in "Geeks". All others are shorter single characters.

Input: s = "abc"
Output: "a"
Explanation: No multi-letter palindromes exist. So the first character "a" is returned as the longest palindromic substring.
 */
func main() {

	s0 := ""
	s1 := "forgeeksskeegfor"
	s2 := "Geeks"
	s3 := "abc"
	s4 := "你好世界🤡🤡界世好你。。"

	fmt.Println(do(s0))
	fmt.Println(do(s1))
	fmt.Println(do(s2))
	fmt.Println(do(s3))
	fmt.Println(do(s4))

}

func do(s string) string {
	slice := []rune(s)
	n := len(slice)

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
		for left >= 0 && right < n && slice[left] == slice[right] {
			// 这里实际上又是摒弃了空隙的单位，因此最终right - left -1的结果不用除以2
			left--
			right++
		}
		// 假设索引间隔n，那么字符串长度则n+1；比如[2, 4]（假设对应字符串为aba），区间长度2，而2，3，4对应元素个数是3。因此+1
		// temp := (right - 1) - (left + 1) + 1
		temp := right - left - 1
		if temp > maxLen {
			maxLen = temp
			// 跳出inner for时，left指针已经往左侧偏移了，因此此处需要+1
			start = left + 1
		}
	}
	return string(slice[start : start+maxLen])
}
