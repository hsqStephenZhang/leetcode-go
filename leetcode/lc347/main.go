package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	key int
	val int
}

type BigHeap []Pair

func (h BigHeap) Len() int           { return len(h) }
func (h BigHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h BigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *BigHeap) Pop() interface{} {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	bigHeap := BigHeap{}
	m := make(map[int]int)
	res := make([]int, k)
	for _, n := range nums {
		m[n]++
	}

	for key, val := range m {
		heap.Push(&bigHeap, Pair{key, val})
	}
	for i := 0; i < k; i++ {
		pair := heap.Pop(&bigHeap).(Pair)
		res[i] = pair.key
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 3, 3}
	k := 1
	fmt.Println((topKFrequent(nums, k)))
}
