package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{}

	cur := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= cur[1] {
			cur[1] = max(cur[1], intervals[i][1])
			cur[0] = min(cur[0], intervals[i][0])
		} else {
			ans = append(ans, cur)
			cur = intervals[i]
		}
	}

	ans = append(ans, cur)

	return ans
}

func main() {
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}
