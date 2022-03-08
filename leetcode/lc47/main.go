package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	used := make([]bool, len(nums))
	ans := [][]int{}
	inner(nums, used, &[]int{}, &ans)
	return ans
}

func inner(nums []int, used []bool, path *[]int, ans *[][]int) {
	if len(*path) == len(nums) {
		newPath := make([]int, len(*path))
		copy(newPath, *path)
		*ans = append(*ans, newPath)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		if !used[i] {
			*path = append(*path, nums[i])
			used[i] = true
			inner(nums, used, path, ans)
			used[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
}

func main() {
	nums := []int{1, 3, 2}
	ans := permuteUnique(nums)
	for _, v := range ans {
		fmt.Println(v)
	}
}
