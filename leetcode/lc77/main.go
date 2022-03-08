package main

import "fmt"

func combine(n int, k int) [][]int {
	ans := [][]int{}
	path := []int{}
	inner(n, 1, k, &path, &ans)
	return ans
}

func inner(n int, index int, k int, path *[]int, ans *[][]int) {
	if k == 0 {
		newPath := make([]int, len(*path))
		copy(newPath, *path)
		*ans = append(*ans, newPath)
		return
	}

	for i := index; i <= n; i++ {
		*path = append(*path, i)
		inner(n, i+1, k-1, path, ans)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}
