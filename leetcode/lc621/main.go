package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func leastInterval(tasks []byte, n int) int {
	length := len(tasks)
	count := make([]int, 26)
	for _, task := range tasks {
		count[task-'A']++
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	m := count[0]
	maxCnt := 1
	for i := 1; i < len(count); i++ {
		if count[i] == m {
			maxCnt++
		} else {
			break
		}
	}

	return max(length, (m-1)*(n+1)+maxCnt)
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
