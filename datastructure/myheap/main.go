package main

import (
	"fmt"
)

func TestHeap() {
	h := BuildHeap([]int{5, 3, 2, 4, 1})
	for i := 0; i < 5; i++ {
		v, _ := h.Pop()
		fmt.Println(v)
	}
}

func main() {
	TestHeap()
}
