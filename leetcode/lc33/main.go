package main

import "fmt"

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1}
	target := 1
	fmt.Println(search(nums, target))
}
