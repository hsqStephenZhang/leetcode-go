package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := [][]int{{}}
	used := make([]bool, len(nums))
	path := []int{}
	inner(nums, used, 0, &ans, &path)
	return ans
}

func inner(nums []int, used []bool, index int, ans *[][]int, path *[]int) {
	if index == len(nums) {
		return
	}
	for i := index; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		used[i] = true
		*path = append(*path, nums[i])
		newPath := make([]int, len(*path))
		copy(newPath, *path)
		*ans = append(*ans, newPath)
		inner(nums, used, i+1, ans, path)
		*path = (*path)[:len(*path)-1]
		used[i] = false
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
