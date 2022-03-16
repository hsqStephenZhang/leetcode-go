package main

import (
	"fmt"
	"math"
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

func networkDelayTime(times [][]int, n int, k int) int {
	const INF = math.MaxInt32
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = INF
		}
	}

	for _, time := range times {
		i, j := time[0]-1, time[1]-1
		graph[i][j] = time[2]
	}

	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = INF
	}
	distance[k-1] = 0

	visited := make([]bool, n)

	for round := 0; round < n; round++ {
		cur := -1
		for i := 0; i < n; i++ {
			if !visited[i] && (cur == -1 || distance[i] < distance[cur]) {
				cur = i
			}
		}
		visited[cur] = true
		for idx, val := range graph[cur] {
			if val != INF {
				distance[idx] = min(distance[idx], distance[cur]+val)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if distance[i] == INF {
			return -1
		}
		res = max(res, distance[i])
	}
	return res
}

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	fmt.Println(networkDelayTime(times, n, k))
}
