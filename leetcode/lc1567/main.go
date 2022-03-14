package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxLen(nums []int) int {
	res := 0
	posDp := make([]int, len(nums)+1)
	negDp := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		num := nums[i-1]
		if num == 0 {
			posDp[i] = 0
			negDp[i] = 0
		} else if num > 0 {
			posDp[i] = posDp[i-1] + 1
			if negDp[i-1] == 0 {
				negDp[i] = 0
			} else {
				negDp[i] = negDp[i-1] + 1
			}
		} else {
			negDp[i] = posDp[i-1] + 1
			if negDp[i-1] == 0 {
				posDp[i] = 0
			} else {
				posDp[i] = negDp[i-1] + 1
			}
		}
		res = max(res, posDp[i])
	}
	return res
}

func main() {
	fmt.Println(getMaxLen([]int{0, 1, -2, -3, -4}))
}
