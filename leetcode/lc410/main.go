package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = int(^uint32((0)) >> 1)
		}
	}
	dp[1][0] = 0
	for i := 1; i <= n; i++ {
		dp[1][i] = dp[1][i-1] + nums[i-1]
	}
	for i := 2; i <= m; i++ {
		for j := i; j <= n; j++ {
			for k := i - 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], max(dp[1][j]-dp[1][k], dp[i-1][k]))
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}
