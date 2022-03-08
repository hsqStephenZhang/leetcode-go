package main

import (
	"fmt"
	"sort"
)

/**
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i int, j int) bool {
		return candidates[i] < candidates[j]
	})
	ans := [][]int{}
	path := []int{}
	used := make([]bool, len(candidates))
	inner(candidates, target, 0, used, &path, &ans)
	return ans
}

func inner(candidates []int, target int, index int, used []bool, path *[]int, ans *[][]int) {
	if target == 0 {
		newPath := make([]int, len(*path))
		copy(newPath, *path)
		*ans = append(*ans, newPath)
		return
	}

	for i := index; i < len(candidates) && candidates[i] <= target; i++ {
		if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
			continue
		}
		*path = append(*path, candidates[i])
		used[i] = true
		inner(candidates, target-candidates[i], i+1, used, path, ans)
		used[i] = false
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	ans := combinationSum2(candidates, target)
	for _, v := range ans {
		fmt.Println(v)
	}
}
