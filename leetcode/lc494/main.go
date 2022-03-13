package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if (sum+target)%2 != 0 {
		return 0
	}

	targetMax := (sum + target) / 2
	length := len(nums)

	dp := make([]int, targetMax+1)
	dp[0] = 1
	for i := 0; i < length; i++ {
		for j := targetMax; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[targetMax]
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}
