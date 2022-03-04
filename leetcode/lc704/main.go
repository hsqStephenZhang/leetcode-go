package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 4, 6}
	target := 5
	fmt.Println(search(nums, target))
}
