package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		n2, n3, n5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(n2, n3), n5)
		// 2*3 == 3*2, so there are some intersections
		if dp[i] == n2 {
			p2++
		}
		if dp[i] == n3 {
			p3++
		}
		if dp[i] == n5 {
			p5++
		}
	}
	return dp[n]
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(nthUglyNumber(i))
	}
	// fmt.Println(nthUglyNumber(10))
}
