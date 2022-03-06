package main

import "fmt"

// 原地修改数组，需要进行标记
func setZeroes(matrix [][]int) {
	numRows := len(matrix)
	numCols := len(matrix[0])
	flag := false
	for i := 0; i < numRows; i++ {
		if matrix[i][0] == 0 {
			flag = true
		}
		for j := 1; j < numCols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := numRows - 1; i >= 0; i-- {
		for j := 1; j < numCols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if flag {
			matrix[i][0] = 0
		}
	}
}

func main() {
	// matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	// setZeroes(matrix)
	// for _, v := range matrix {
	// 	fmt.Println(v)
	// }

	matrix2 := [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}
	setZeroes(matrix2)
	for _, v := range matrix2 {
		fmt.Println(v)
	}

	// matrix := [][]int{{0}, {1}}
	// setZeroes(matrix)
	// for _, v := range matrix {
	// 	fmt.Println(v)
	// }
}
