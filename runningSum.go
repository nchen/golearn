// https://leetcode-cn.com/problems/running-sum-of-1d-array/submissions/
package main

func runningSum(nums []int) []int {
	// 声明并分配一个动态数组
	var result = make([]int, len(nums))

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result[i] = sum
	}
	return result
}
