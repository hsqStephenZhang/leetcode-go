package main

import (
	"fmt"
)

type Heap struct {
	data []int
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func BuildHeap(data []int) *Heap {
	h := &Heap{data}
	n := h.Len()
	for innerNode := n/2 - 1; innerNode >= 0; innerNode-- {
		down(h, innerNode, n)
	}
	return h
}

func (h *Heap) Push(v int) {
	h.data = append(h.data, v)
	up(h, h.Len()-1)
}

func (h *Heap) Pop() (int, error) {
	if len(h.data) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	n := len(h.data) - 1
	h.Swap(0, n)
	down(h, 0, n)
	v := h.data[n]
	h.data = h.data[:n]
	return v, nil
}

func down(h *Heap, parent0 int, n int) bool {
	parent := parent0
	for {
		leftChild := 2*parent + 1
		if leftChild >= n || leftChild < 0 {
			break
		}
		choosed := leftChild
		if rightChild := leftChild + 1; rightChild < n && h.Less(rightChild, leftChild) {
			choosed = rightChild
		}
		if !h.Less(choosed, parent) {
			break
		}
		h.Swap(parent, choosed)
		parent = choosed
	}
	return parent > parent0
}

func up(h *Heap, i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || !h.Less(i, parent) {
			break
		}
		h.Swap(parent, i)
		i = parent
	}
}
