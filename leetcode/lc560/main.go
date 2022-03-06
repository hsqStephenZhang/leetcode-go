package main

import "fmt"

// 整数数组，不是正整数数组
// func subarraySum(nums []int, k int) int {
// 	length := len(nums)
// 	right := 0
// 	left := 0
// 	sum := 0
// 	res := 0
// 	for right < length {
// 		sum += nums[right]
// 		if sum == k {
// 			res += 1
// 		} else if sum > k {
// 			for left < right {
// 				sum -= nums[left]
// 				left += 1
// 				if sum == k {
// 					res += 1
// 					break
// 				} else if sum < k {
// 					break
// 				}
// 			}
// 		}
// 		right += 1
// 	}
// 	return res
// }

func subarraySum(nums []int, k int) int {

	return 1
}

func main() {
	nums := []int{-1, -1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}
