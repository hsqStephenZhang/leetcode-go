package main

import (
	"fmt"
)

func canCross(stones []int) bool {
	if len(stones) == 1 {
		return true
	}
	val2idx := make(map[int]int)
	n := len(stones)
	// dp[i][j] 表示跳到了 i, 且最后一次跳跃距离为 j，
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		val2idx[stones[i]] = i
	}

	if stones[1]-stones[0] == 1 {
		dp[1][1] = true
	} else {
		return false
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] {
				if idx1, ok1 := val2idx[stones[i]+j-1]; ok1 {
					dp[idx1][j-1] = true
				}
				if idx1, ok1 := val2idx[stones[i]+j]; ok1 {
					dp[idx1][j] = true
				}
				if idx1, ok1 := val2idx[stones[i]+j+1]; ok1 {
					dp[idx1][j+1] = true
				}
			}
		}
	}

	for j := 1; j < n; j++ {
		if dp[n-1][j] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	fmt.Println(canCross([]int{0, 2}))
}
