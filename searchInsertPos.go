// https://leetcode-cn.com/problems/search-insert-position/
package main

func searchInsert(nums []int, target int) int {
	result := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			result = i
			break
		}
	}
	return result
}
