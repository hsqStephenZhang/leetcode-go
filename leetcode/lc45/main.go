package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	numSteps := 0

	length := len(nums)
	curPos := 0
	nextPos := 0

	for i := 0; i < length; i++ {
		if i > curPos {
			numSteps++
			curPos = nextPos
		}
		nextPos = max(nextPos, nums[i]+i)
	}

	return numSteps
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
