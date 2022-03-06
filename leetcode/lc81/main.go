package main

import "fmt"

func search(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[low] && nums[mid] == nums[high] {
			low += 1
			high -= 1
			continue
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
	return false
}

func main() {
	nums := []int{1, 0, 1, 1, 1}
	target := 0
	fmt.Println(search(nums, target))
}
