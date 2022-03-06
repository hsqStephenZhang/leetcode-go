package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})

	ans := [][]int{}
	i := 0
	for i < len(nums)-2 {
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right {
					if nums[left] == nums[left+1] {
						left++
					} else {
						left++
						break
					}
				}
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
		for i < len(nums)-2 {
			if nums[i] == nums[i+1] {
				i++
			} else {
				i++
				break
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	for _, v := range ans {
		fmt.Println(v)
	}
}
