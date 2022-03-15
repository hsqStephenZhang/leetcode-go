package main

import "fmt"

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2

		// [mid, high] is sorted
		if nums[mid] < nums[high] {
			high = mid
		} else { // [low, mid] is sorted
			low = mid + 1
		}
	}
	return nums[low]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
