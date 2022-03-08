package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	num   int
	count int
}

func combinationSum(candidates []int, target int) [][]int {

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	ans := [][]int{}
	path := []Pair{}
	inner(candidates, target, 0, &path, &ans)
	return ans
}

// suitable for the situation where one number may occur more than once
func inner(candidates []int, target int, index int, path *[]Pair, ans *[][]int) {
	if target == 0 {
		// fmt.Println(*path)
		size := 0
		for _, v := range *path {
			size += v.count
		}
		newPath := make([]int, 0, size)
		for _, v := range *path {
			for i := 0; i < v.count; i++ {
				newPath = append(newPath, v.num)
			}
		}
		*ans = append(*ans, newPath)
		return
	}

	if index >= len(candidates) {
		return
	}

	numDiv := target / candidates[index]
	for j := 0; j <= numDiv; j++ {
		*path = append(*path, Pair{candidates[index], j})
		inner(candidates, target-candidates[index]*j, index+1, path, ans)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	ans := combinationSum(candidates, target)
	for _, v := range ans {
		fmt.Println(v)
	}
}
