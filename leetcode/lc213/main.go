package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	return max(inner(nums[1:]), inner(nums[:length-1]))
}

func inner(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	// dp[i] means the max money we can get in nums[0:i]
	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[length-1]
}

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}
