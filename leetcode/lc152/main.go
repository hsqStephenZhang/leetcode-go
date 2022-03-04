package main

import "fmt"

func findMin(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	low := 0
	high := length - 1
	for low <= high {
		// mid := low + (high-low)/2
		// if mid==0 || nums[mid]
	}
	return -1
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
