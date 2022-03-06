package main

import (
	"container/heap"
	"fmt"
)

type SmallHeap []int

func (h SmallHeap) Len() int { return len(h) }

func (h SmallHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h SmallHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SmallHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *SmallHeap) Pop() interface{} {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func findKthLargest(nums []int, k int) int {
	h := SmallHeap(nums[0:k])
	heap.Init(&h)
	for i := k; i < len(nums); i++ {
		if nums[i] > h[0] {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}
	return h[0]
}

type BigHeap []int

func (h BigHeap) Len() int { return len(h) }

func (h BigHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h BigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *BigHeap) Pop() interface{} {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func findKthLargest2(nums []int, k int) int {
	h := BigHeap(nums)
	heap.Init(&h)
	for k > 1 {
		heap.Pop(&h)
		k--
	}
	return heap.Pop(&h).(int)
}

func main() {
	k := 3
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, k))
	fmt.Println(findKthLargest2([]int{3, 2, 1, 5, 6, 4}, k))
}
