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

// left bound search, close interval
func searchLeftBound(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// right bound search
func searchRightBound(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target { // must increase low, otherwise may enter endless loop
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}

			low = mid + 1
		}
	}
	return -1
}

// left bound search, close interval
func searchFirstBigEqual(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// left bound search, close interval
func searchLastLessEqual(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 4, 6}
	target := 5
	fmt.Println(search(nums, target))

	nums = []int{0, 1, 1, 1, 2, 2, 2, 2, 3}
	fmt.Println(searchLeftBound(nums, 1))
	fmt.Println(searchRightBound(nums, 1))

	fmt.Println(searchFirstBigEqual(nums, 2))
	fmt.Println(searchLastLessEqual(nums, 3))
}
