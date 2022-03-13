package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sum := 0
	for _, v := range stones {
		sum += v
	}
	length := len(stones)
	target := sum / 2

	// dp[i][j] means whether any nums[0:i]'s subset can sum to j
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	dp[0][stones[0]] = true
	for i := 1; i < length; i++ {
		dp[i][0] = true
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j] || (j >= stones[i] && dp[i-1][j-stones[i]])
		}
	}
	for i := target; i >= 0; i-- {
		if dp[length-1][i] {
			return sum - 2*i
		}
	}
	// unreachable
	return 0
}

func main() {
	var stones []int
	// stones = []int{2, 7, 4, 1, 8, 1}
	// fmt.Println(lastStoneWeightII(stones))
	// stones = []int{31, 26, 33, 21, 40}
	// fmt.Println(lastStoneWeightII(stones))
	stones = []int{91}
	fmt.Println(lastStoneWeightII(stones))
}
