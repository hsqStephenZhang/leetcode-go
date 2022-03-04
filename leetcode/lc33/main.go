package main

import "fmt"

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[0] <= nums[mid] { // [low, mid) 有序
			if nums[0] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // [mid, high) 有序
			if nums[mid] < target && target <= nums[length-1] {
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
