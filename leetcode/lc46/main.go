package main

import "fmt"

func permute(nums []int) [][]int {
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
	nums := []int{1, 2, 3}
	ans := permute(nums)
	for _, v := range ans {
		fmt.Println(v)
	}
}
