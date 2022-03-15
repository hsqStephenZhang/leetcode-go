package main

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				res += 1
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	steps := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, step := range steps {
		x, y := i+step[0], j+step[1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
			grid[x][y] = '0'
			dfs(grid, x, y)
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '1'},
		{'0', '0', '0', '1', '0'},
	}
	fmt.Println(numIslands(grid))
}
