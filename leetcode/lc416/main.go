package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	length := len(nums)
	target := sum / 2

	// dp[i][j] means whether any nums[0:i]'s subset can sum to j
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	for i := 1; i < length; i++ {
		dp[i][0] = true
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j] || (j >= nums[i] && dp[i-1][j-nums[i]])
		}
	}

	return dp[length-1][target]
}

func main() {
	nums := []int{1, 2, 3, 6}
	fmt.Println(canPartition(nums))
}
