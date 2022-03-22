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

func main() {
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{3, 3, 5, 6, 2, 3}))
}
