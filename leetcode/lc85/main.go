package main

import "fmt"

type Pair struct {
	left int
	idx  int
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []Pair{{left: 0, idx: 0}}

	res := 0

	for i := 1; i < n; i++ {
		left := i
		for len(stack) > 0 && heights[stack[len(stack)-1].idx] >= heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left = top.left
			area := heights[top.idx] * (i - top.left)
			if area > res {
				res = area
			}
		}
		stack = append(stack, Pair{left: left, idx: i})
	}

	for elem := range stack {
		area := heights[stack[elem].idx] * (n - stack[elem].left)
		if area > res {
			res = area
		}
	}

	return res
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	heightMatrix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		heightMatrix[i] = make([]int, n)
	}
	for i := 1; i <= m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i-1][j] == '1' {
				heightMatrix[i][j] = heightMatrix[i-1][j] + 1
			} else {
				heightMatrix[i][j] = 0
			}
		}
	}

	res := 0

	for i := 1; i <= m; i++ {
		r := largestRectangleArea(heightMatrix[i])
		if r > res {
			res = r
		}
	}

	return res
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
