package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	length := len(nums)
	right := length - 1
	for right >= 1 && nums[right-1] >= nums[right] {
		right -= 1
	}
	if right == 0 {
		// reverse the whole array
		for i := 0; i < length/2; i++ {
			nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
		}
		return
	}
	left := right - 1
	idx := length - 1
	for idx >= right && nums[idx] <= nums[left] {
		idx--
	}

	nums[left], nums[idx] = nums[idx], nums[left]
	// reverse nums[right:]
	for i := 0; i < (length-right)/2; i++ {
		nums[i+right], nums[length-1-i] = nums[length-1-i], nums[i+right]
	}
}

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
