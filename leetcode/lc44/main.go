package main

import "fmt"

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for j, c2 := range p {
		if c2 == '*' {
			idx := 0
			for ; idx <= m; idx++ {
				if dp[idx][j] {
					break
				}
			}
			if idx <= m {
				for i := idx; i <= m; i++ {
					dp[i][j+1] = true
				}
			}
		} else {
			for i, c1 := range s {
				if c1 == c2 || c2 == '?' {
					dp[i+1][j+1] = dp[i][j]
				}
			}
		}

	}

	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("aac", "*c"))
}
