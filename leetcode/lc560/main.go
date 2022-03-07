package main

import "fmt"

// 遍历过程当中可以保证前后顺序
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	sum := 0
	m[0] = 1
	ans := 0
	for _, v := range nums {
		sum += v
		if _, ok := m[sum-k]; ok {
			ans += m[sum-k]
		}
		m[sum]++
	}
	return ans
}

func main() {
	nums := []int{-1, -1, 1}
	k := 1
	fmt.Println(subarraySum(nums, k))
}
