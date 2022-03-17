package main

import (
	"container/heap"
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

type Pair struct {
	dist, idx int
}

type SmallHeap []Pair

func (h SmallHeap) Len() int {
	return len(h)
}

func (h SmallHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h SmallHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SmallHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *SmallHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	const INF = math.MaxInt32
	adjacent := make(map[int][]Pair)

	for _, time := range times {
		i, j := time[0]-1, time[1]-1
		adjacent[i] = append(adjacent[i], Pair{time[2], j})
	}

	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = INF
	}
	distance[k-1] = 0

	h := &SmallHeap{}

	visited := make([]bool, n)

	heap.Push(h, Pair{0, k - 1})

	for h.Len() > 0 {
		cur := heap.Pop(h).(Pair)
		if visited[cur.idx] {
			continue
		}
		visited[cur.idx] = true
		for _, p := range adjacent[cur.idx] {
			d, j := p.dist, p.idx
			if distance[cur.idx]+d < distance[j] {
				distance[j] = distance[cur.idx] + d
				heap.Push(h, Pair{distance[j], j})
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

func networkDelayTime2(times [][]int, n int, k int) int {
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
	times := [][]int{{1, 2, 1}}
	n := 2
	k := 2
	fmt.Println(networkDelayTime(times, n, k))
	fmt.Println(networkDelayTime2(times, n, k))
}
