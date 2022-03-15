package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, index int, word string) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[index] || board[i][j] == '*' {
		return false
	}
	tmp := board[i][j]
	board[i][j] = '#'
	if dfs(board, i+1, j, index+1, word) ||
		dfs(board, i-1, j, index+1, word) ||
		dfs(board, i, j+1, index+1, word) ||
		dfs(board, i, j-1, index+1, word) {
		return true
	}
	board[i][j] = tmp
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCB"
	fmt.Println(exist(board, word))
}
