package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	rows := 0
	for i := n - 1; i > 0; i -= 2 {
		for offset := 0; offset < i; offset++ {
			matrix[rows][rows+offset], matrix[rows+offset][n-1-rows], matrix[n-1-rows][n-1-rows-offset], matrix[n-1-rows-offset][rows] = matrix[n-1-rows-offset][rows], matrix[rows][rows+offset], matrix[rows+offset][n-1-rows], matrix[n-1-rows][n-1-rows-offset]
		}
		rows += 1
	}
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	for _, v := range matrix {
		fmt.Println(v)
	}
}
