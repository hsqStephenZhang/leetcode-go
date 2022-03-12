package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	key   int
	value int
	cnt   int
}

type LFUCache struct {
	freq2list map[int]*list.List
	key2node  map[int]*list.Element
	capacity  int
	minFreq   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freq2list: make(map[int]*list.List),
		key2node:  make(map[int]*list.Element),
		capacity:  capacity,
		minFreq:   0,
	}
}

func (cache *LFUCache) Get(key int) int {
	elem, ok := cache.key2node[key]
	if !ok {
		return -1
	}
	node := elem.Value.(*Node)
	oldCnt := node.cnt
	cache.freq2list[oldCnt].Remove(elem)
	node.cnt++
	newCnt := oldCnt + 1
	if _, ok := cache.freq2list[newCnt]; !ok {
		cache.freq2list[newCnt] = list.New()
	}
	newNode := cache.freq2list[newCnt].PushFront(node)
	cache.key2node[key] = newNode
	if cache.minFreq == oldCnt && cache.freq2list[oldCnt].Len() == 0 {
		cache.minFreq++
	}
	return node.value
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}
	if elem, ok := cache.key2node[key]; ok {
		node := elem.Value.(*Node)
		node.value = value
		cache.Get(key)
		return
	}

	if len(cache.key2node) == cache.capacity {
		elem := cache.freq2list[cache.minFreq].Back()
		node := elem.Value.(*Node)
		cache.freq2list[cache.minFreq].Remove(elem)
		delete(cache.key2node, node.key)
	}
	cache.minFreq = 1
	newNode := &Node{key: key, value: value, cnt: 1}
	if _, ok := cache.freq2list[1]; !ok {
		cache.freq2list[1] = list.New()
	}
	newElem := cache.freq2list[1].PushFront(newNode)
	cache.key2node[key] = newElem
}

func main() {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Get(2)
	cache.Put(1, 1)
	cache.Put(4, 1)
	fmt.Println(cache.Get(2))

	// fmt.Println(cache.Get(1))
	// fmt.Println(cache.Get(3))
	// fmt.Println(cache.Get(4))
	// fmt.Println(cache.Get(5))
}
