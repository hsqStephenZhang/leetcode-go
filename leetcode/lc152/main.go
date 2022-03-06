package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMin(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	low := 0
	high := length - 1
	minVal := 100000000
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= nums[low] {
			minVal = min(minVal, nums[low])
			low = mid + 1
		} else {
			high = mid - 1
			minVal = min(minVal, nums[mid])
		}
	}
	return minVal
}

func main() {
	nums := []int{3, 1, 2}
	fmt.Println(findMin(nums))
}
