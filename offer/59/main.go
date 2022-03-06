package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	length := len(nums)
	l := make([]int, 0, length)
	res := make([]int, length-k+1)

	for i := 0; i < k; i++ {
		if len(l) == 0 {
			l = append(l, i)
			continue
		}
		cur := nums[i]
		for len(l) > 0 && nums[l[len(l)-1]] < cur {
			l = l[:len(l)-1]
		}
		l = append(l, i)
	}

	res[0] = nums[l[0]]

	for i := k; i < length; i++ {
		if nums[l[0]] == nums[i-k] {
			l = l[1:]
		}
		cur := nums[i]
		for len(l) > 0 && nums[l[len(l)-1]] < cur {
			l = l[:len(l)-1]
		}
		l = append(l, i)
		res[i-k+1] = nums[l[0]]
	}

	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
