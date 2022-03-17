package main

import (
	"container/heap"
	"fmt"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pos struct {
	x, y int
}

type Pair struct {
	dist int
	pos  Pos
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

func minimumEffortPath(heights [][]int) int {
	const INF = math.MaxInt32
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	n := len(heights)
	m := len(heights[0])

	h := &SmallHeap{}
	heap.Push(h, Pair{0, Pos{0, 0}})

	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distance[i][j] = INF
		}
	}
	distance[0][0] = 0

	visited := make(map[Pos]bool)
	for h.Len() > 0 {
		cur := heap.Pop(h).(Pair)
		if cur.pos.x == n-1 && cur.pos.y == m-1 {
			return cur.dist
		}
		if visited[cur.pos] {
			continue
		}
		visited[cur.pos] = true
		// traverse adjacent position
		for _, dir := range dirs {
			x, y := cur.pos.x+dir[0], cur.pos.y+dir[1]
			if x < 0 || y < 0 || x >= n || y >= m {
				continue
			}
			dis1 := distance[cur.pos.x][cur.pos.y]
			dis2 := distance[x][y]
			dis := max(dis1, abs(heights[x][y]-heights[cur.pos.x][cur.pos.y]))
			if dis < dis2 {
				distance[x][y] = dis
				heap.Push(h, Pair{dis, Pos{x, y}})
			}
		}
	}

	// unreachable
	return -1
}

func main() {
	var height [][]int

	height = [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(height))

	height = [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}
	fmt.Println(minimumEffortPath(height))

	height = [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(height))
}
