package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := [][]int{{}}
	for _, v := range nums {
		length := len(ans)
		for i := 0; i < length; i++ {
			ans = append(ans, append([]int{v}, ans[i]...))
		}
	}
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}
