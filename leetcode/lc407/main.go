package main

import (
	"container/heap"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pair struct {
	dist, x, y int
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

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	n := len(heightMap[0])
	h := &SmallHeap{}

	heap.Init(h)

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		heap.Push(h, Pair{heightMap[i][0], i, 0})
		heap.Push(h, Pair{heightMap[i][n-1], i, n - 1})
		visited[i][0] = true
		visited[i][n-1] = true
	}

	for j := 1; j < n-1; j++ {
		heap.Push(h, Pair{heightMap[0][j], 0, j})
		heap.Push(h, Pair{heightMap[m-1][j], m - 1, j})
		visited[0][j] = true
		visited[m-1][j] = true
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	ans := 0

	for h.Len() > 0 {
		cur := heap.Pop(h).(Pair)
		for _, dir := range dirs {
			x := cur.x + dir[0]
			y := cur.y + dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
				continue
			}
			if cur.dist > heightMap[x][y] {
				ans += cur.dist - heightMap[x][y]
			}
			// 想象上一步之后，将 [x][y] 用水填充到 cur.dist 高度
			// 这样， cur, [x][y] 之间高度相同，相当于将边界向内扩充了
			heap.Push(h, Pair{max(cur.dist, heightMap[x][y]), x, y})
			visited[x][y] = true
		}
	}

	return ans
}

func main() {
	var heightMap [][]int
	heightMap = [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	fmt.Println(trapRainWater(heightMap))

	heightMap = [][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}

	fmt.Println(trapRainWater(heightMap))
}
