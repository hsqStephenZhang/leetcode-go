package main

import (
	"container/heap"
	"fmt"
)

type KthLargest struct {
	inner SmallHeap
	k     int
}

type SmallHeap []int

func (h SmallHeap) Len() int {
	return len(h)
}

func (h SmallHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h SmallHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SmallHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *SmallHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	h := SmallHeap{}
	if len(nums) < k {
		for _, v := range nums {
			heap.Push(&h, v)
		}

	} else {
		for _, v := range nums[0:k] {
			heap.Push(&h, v)
		}
		for _, v := range nums[k:] {
			if v > h[0] {
				heap.Pop(&h)
				heap.Push(&h, v)
			}
		}
	}

	return KthLargest{h, k}
}

func (me *KthLargest) Add(val int) int {
	if len(me.inner) < me.k {
		heap.Push(&me.inner, val)
	} else if val > me.inner[0] {
		heap.Pop(&me.inner)
		heap.Push(&me.inner, val)
	}
	return me.inner[0]
}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, nums)
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))
}
