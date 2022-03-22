package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			// 0 表示不持有股票，1 表示持有股票
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i <= k; i++ {
		dp[0][i][1] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
}

func main() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
