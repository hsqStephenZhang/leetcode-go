package main

import "fmt"

type State struct {
	row, col int
}

func exist(board [][]byte, word string) bool {
	memory := make(map[State]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			val, ok := memory[State{i, j}]
			if !ok || !val {
				if dfs(board, i, j, 0, word, memory) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, index int, word string, memory map[State]bool) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[index] || memory[State{i, j}] {
		return false
	}
	memory[State{i, j}] = true
	if dfs(board, i+1, j, index+1, word, memory) ||
		dfs(board, i-1, j, index+1, word, memory) ||
		dfs(board, i, j+1, index+1, word, memory) ||
		dfs(board, i, j-1, index+1, word, memory) {
		return true
	}
	memory[State{i, j}] = false
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
