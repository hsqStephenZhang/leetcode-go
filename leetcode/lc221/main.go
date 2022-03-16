package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	maxWidth := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(min(dp[i][j+1], dp[i+1][j]), dp[i][j]) + 1
			}
			if dp[i+1][j+1] > maxWidth {
				maxWidth = dp[i+1][j+1]
			}
		}
	}
	return maxWidth * maxWidth
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},
	}
	fmt.Println(maximalSquare(matrix))
}
