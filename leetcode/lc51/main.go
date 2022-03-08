package main

import "fmt"

// usedCols[i] 表示第 i 列在第 x 行被占用
func solveNQueens(n int) [][]string {
	ans := [][]string{}
	usedCols := make([]int, n)
	for i := 0; i < n; i++ {
		usedCols[i] = -1
	}
	inner(n, 0, usedCols, &ans)
	return ans
}

func inner(n int, rowIndex int, usedCols []int, ans *[][]string) {
	if rowIndex == n {
		// push the answer
		// fmt.Println(usedCols)
		board := make([]string, n)
		for i := 0; i < n; i++ {
			board[i] = ""
			for j := 0; j < n; j++ {
				if usedCols[j] == i {
					board[i] += "Q"
				} else {
					board[i] += "."
				}
			}
		}

		*ans = append(*ans, board)
		return
	}

	for colIndex := 0; colIndex < n; colIndex++ {
		if usedCols[colIndex] == -1 && isValid(usedCols, rowIndex, colIndex) {
			usedCols[colIndex] = rowIndex
			inner(n, rowIndex+1, usedCols, ans)
			usedCols[colIndex] = -1
		}
	}
}

func isValid(used []int, rowIndex0, colIndex0 int) bool {
	n := len(used)
	rowIndex, colIndex := rowIndex0, colIndex0
	for ; rowIndex >= 0 && colIndex >= 0; rowIndex, colIndex = rowIndex-1, colIndex-1 {
		if used[colIndex] == rowIndex {
			return false
		}
	}
	rowIndex, colIndex = rowIndex0, colIndex0
	for ; rowIndex >= 0 && colIndex < n; rowIndex, colIndex = rowIndex-1, colIndex+1 {
		if used[colIndex] == rowIndex {
			return false
		}
	}
	return true
}

func main() {
	n := 8
	ans := solveNQueens(n)
	fmt.Println(len(ans))
}
