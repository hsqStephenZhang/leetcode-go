package main

import "fmt"

func lengthOfLIS(nums []int) int {
	res := 0
	stack := []int{}
	for _, num := range nums {
		if len(stack) == 0 || stack[len(stack)-1] < num {
			stack = append(stack, num)
			res++
		} else {
			index := binarySearch(stack, num)
			if index >= 0 {
				stack[index] = num
			}
		}
		fmt.Println(stack)
	}
	return res
}

// return index of the first element in nums that is greater or equal to target
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1
		}
	}
	// left = right + 1
	if left >= len(nums) || nums[left] < target {
		return -1
	}
	return left
}

func main() {
	nums := []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}
	fmt.Println(lengthOfLIS(nums))
}
