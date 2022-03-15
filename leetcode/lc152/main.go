package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 以 nums[i] 结尾的乘积绝对值最大的子序列乘积
	posDp := make([]int, len(nums)+1)
	negDp := make([]int, len(nums)+1)
	posDp[0] = 0
	negDp[0] = 0

	maxP := 0

	for idx, num := range nums {
		if num == 0 {
			posDp[idx+1] = 0
			negDp[idx+1] = 0
		} else if num > 0 {
			if posDp[idx] == 0 {
				posDp[idx+1] = num
			} else {
				posDp[idx+1] = num * posDp[idx]
			}
			if negDp[idx] == 0 {
				negDp[idx+1] = 0
			} else {
				negDp[idx+1] = negDp[idx] * num
			}
		} else {
			if posDp[idx] == 0 {
				negDp[idx+1] = num
			} else {
				negDp[idx+1] = num * posDp[idx]
			}
			if negDp[idx] == 0 {
				posDp[idx+1] = 0
			} else {
				posDp[idx+1] = negDp[idx] * num
			}
		}
		if posDp[idx+1] > maxP {
			maxP = posDp[idx+1]
		}
	}
	return maxP
}

func main() {
	fmt.Println(maxProduct([]int{2, 0, -2, 4, -2}))
}
