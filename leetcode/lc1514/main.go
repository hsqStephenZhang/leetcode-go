package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	dist  float64
	index int
}

type BigHeap []Pair

func (h BigHeap) Len() int {
	return len(h)
}

func (h BigHeap) Less(i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h BigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *BigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	h := &BigHeap{}
	heap.Push(h, Pair{1.0, start})

	distance := make([]float64, n)
	for i := 0; i < n; i++ {
		distance[i] = 0.00000001
	}
	distance[start] = 1.0

	adjacent := make(map[int][]Pair)
	for idx, edge := range edges {
		adjacent[edge[0]] = append(adjacent[edge[0]], Pair{succProb[idx], edge[1]})
		adjacent[edge[1]] = append(adjacent[edge[1]], Pair{succProb[idx], edge[0]})
	}

	visited := make(map[int]bool)
	for h.Len() > 0 {
		cur := heap.Pop(h).(Pair)
		if cur.index == end {
			return distance[cur.index]
		}
		if visited[cur.index] {
			continue
		}
		visited[cur.index] = true
		// traverse adjacent position
		for _, pair := range adjacent[cur.index] {
			if distance[pair.index] < distance[cur.index]*pair.dist {
				distance[pair.index] = distance[cur.index] * pair.dist
				heap.Push(h, Pair{distance[pair.index], pair.index})
			}
		}
	}

	// unreachable
	return 0
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2
	fmt.Println(maxProbability(n, edges, succProb, start, end))

}
