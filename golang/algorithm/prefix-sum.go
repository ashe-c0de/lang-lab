package main

import (
	"fmt"
)

/*
 */
func main() {
	nums0 := []int{1, 2, 3}
	nums1 := []int{1, 0, 3}
	nums2 := []int{1, 1, 2, -1, 3}
	fmt.Println(Execute(nums0, 3))
	fmt.Println(Execute(nums1, 1))
	fmt.Println(Execute(nums2, 2))
}

func Execute(nums []int, k int) int {
	res, pre := 0, 0
	m := make(map[int]int, len(nums))
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := m[pre-k]; ok {
			res += v
		}
		m[pre] += 1

	}
	return res
}
