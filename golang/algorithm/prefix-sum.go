package main

import (
	"fmt"
)

/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:
Input: nums = [1,1,1], k = 2

Output: 2

Example 2:
Input: nums = [1,2,3], k = 3

Output: 2 
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
		// 计算前缀和
		pre += nums[i]
		if v, ok := m[pre-k]; ok {
			res += v
		}
		// 存入哈希表，vlue为前缀和key所出现的次数
		m[pre] += 1
	}
	return res
}
