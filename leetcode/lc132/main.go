package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = n
		for j := i - 1; j >= 0; j-- {
			if isPalindrome(s[j:i]) {
				if j == 0 {
					dp[i] = 0
				} else {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(minCut("aabcbac"))
}
