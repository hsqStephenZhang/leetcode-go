package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	n := len(ratings)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			// dp[i] 是从左到右推导出来的
			// dp[i+1]+1 是从右到左推导出来的
			dp[i] = max(dp[i], dp[i+1]+1)
		}
	}
	sum := 0
	for _, v := range dp {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}
