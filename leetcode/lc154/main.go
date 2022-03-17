package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMin(nums []int) int {

	length := len(nums)
	left := 0
	right := length - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] == nums[right] {
			right -= 1
		} else if nums[mid] > nums[left] {
			left = mid + 1
		} else if nums[mid] == nums[left] {
			left += 1
		}
	}

	return min(nums[left], nums[right])
}

func main() {
	nums := []int{2, 2, 2, -2, -1, 1, 1}
	fmt.Println(findMin(nums))
}
