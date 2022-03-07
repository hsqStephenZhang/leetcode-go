package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := [][]int{}
	length := len(nums)
	first := 0
	for first < length-3 {
		second := first + 1
		for second < length-2 {
			onetwosum := target - (nums[first] + nums[second])
			left := second + 1
			right := length - 1
			for left < right {
				sum := nums[left] + nums[right]
				if sum == onetwosum {
					ans = append(ans, []int{nums[first], nums[second], nums[left], nums[right]})
					for left < right {
						if nums[left] == nums[left+1] {
							left++
						} else {
							left++
							break
						}
					}
				} else if sum < onetwosum {
					left++
				} else {
					right--
				}
			}

			for second < len(nums)-2 {
				if nums[second] == nums[second+1] {
					second++
				} else {
					second++
					break
				}
			}
		}

		for first < len(nums)-3 {
			if nums[first] == nums[first+1] {
				first++
			} else {
				first++
				break
			}
		}

	}

	return ans
}

func main() {
	// nums := []int{-2, -1, 0, 0, 1, 2}
	// target := 0
	// ans := fourSum(nums, target)
	// for _, v := range ans {
	// 	fmt.Println(v)
	// }

	nums := []int{2, 2, 2, 2, 2}
	target := 8
	ans := fourSum(nums, target)
	for _, v := range ans {
		fmt.Println(v)
	}
}
