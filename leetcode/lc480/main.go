package main

import (
	"container/heap"
	"fmt"
)

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

type BigHeap []int

func (h BigHeap) Len() int {
	return len(h)
}

func (h BigHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h BigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *BigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type DoubleHeap struct {
	small         SmallHeap
	big           BigHeap
	delay         map[int]struct{}
	smallDelayNum int
	bigDelayNum   int
}

func NewDoubleHeap() *DoubleHeap {
	h := &DoubleHeap{
		small:         SmallHeap{},
		big:           BigHeap{},
		delay:         map[int]struct{}{},
		smallDelayNum: 0,
		bigDelayNum:   0,
	}
	return h
}

func (h *DoubleHeap) Push(cur int) {
	for {
		if len(h.small) == 0 || len(h.big) == 0 {
			break
		}
		_, ok1 := h.delay[h.small[0]]
		_, ok2 := h.delay[h.big[0]]
		if ok1 && ok2 {
			heap.Pop(&h.small)
			heap.Pop(&h.big)
			delete(h.delay, h.small[0])
			delete(h.delay, h.big[0])
			h.smallDelayNum--
			h.bigDelayNum--
		} else if ok1 && !ok2 {
			if h.small.Len()-h.smallDelayNum+1 > h.big.Len()-h.bigDelayNum {
				val := heap.Pop(&h.small).(int)
				delete(h.delay, val)
				h.smallDelayNum--
			} else {
				val := heap.Pop(&h.small).(int)
				delete(h.delay, val)
				h.smallDelayNum--
				another := heap.Pop(&h.big)
				heap.Push(&h.small, another)
			}
		} else if !ok1 && ok2 {
			if h.small.Len()-h.smallDelayNum-1 > h.big.Len()-h.bigDelayNum {
				val := heap.Pop(&h.big).(int)
				delete(h.delay, val)
				h.bigDelayNum--
				another := heap.Pop(&h.small)
				heap.Push(&h.big, another)
			} else {
				val := heap.Pop(&h.big).(int)
				delete(h.delay, val)
				h.bigDelayNum--
			}
		} else {
			break
		}
	}
	if h.small.Len() != 0 && h.big.Len() == 0 {
		heap.Push(&h.big, cur)
		return
	} else if h.big.Len() == 0 && h.small.Len() == 0 {
		heap.Push(&h.small, cur)
		return
	}

	midBigger := h.small[0]
	midSmaller := h.big[0]

	if h.small.Len()-h.smallDelayNum > h.big.Len()-h.bigDelayNum {
		if cur <= midBigger {
			heap.Push(&h.small, cur)
		} else {
			heap.Pop(&h.small)
			heap.Push(&h.small, cur)
			heap.Push(&h.big, midBigger)
		}
	} else { // equal
		if cur >= midSmaller {
			heap.Push(&h.small, cur)
		} else {
			heap.Pop(&h.big)
			heap.Push(&h.small, midSmaller)
			heap.Push(&h.big, cur)
		}
	}
}

func (h *DoubleHeap) Remove(val int) {
	midBigger := h.small[0]
	h.delay[val] = struct{}{}
	if val >= midBigger {
		h.smallDelayNum++
	} else {
		h.bigDelayNum++
	}
}

func (h *DoubleHeap) Mid() float64 {
	if h.small.Len() > h.big.Len() {
		return float64(h.small[0])
	} else {
		return (float64(h.big[0]) + float64(h.small[0])) / 2
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	if k == 1 {
		ans := make([]float64, len(nums))
		for i := 0; i < len(nums); i++ {
			ans[i] = float64(nums[i])
		}
		return ans
	}
	h := NewDoubleHeap()
	res := []float64{}

	for i := 0; i < k; i++ {
		h.Push(nums[i])
	}
	res = append(res, h.Mid())
	for i := k; i < len(nums); i++ {
		h.Remove(nums[i-k])
		h.Push(nums[i])
		res = append(res, h.Mid())
	}

	return res
}

func main() {
	nums := []int{2, 3, 4, 5, 6}
	k := 2
	ans := medianSlidingWindow(nums, k)
	fmt.Println(ans)
}
