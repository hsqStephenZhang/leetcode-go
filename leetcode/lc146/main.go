package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	key   int
	value int
}

type LRUCache struct {
	key2node map[int]*list.Element
	data     list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		key2node: make(map[int]*list.Element),
		data:     list.List{},
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	elem, ok := cache.key2node[key]
	if !ok {
		return -1
	}
	cache.data.MoveToFront(elem)
	return elem.Value.(*Node).value
}

func (cache *LRUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}
	// case 1
	if elem, ok := cache.key2node[key]; ok {
		cache.data.MoveToFront(elem)
		node := elem.Value.(*Node)
		node.value = value
		return
	}
	// case 2, delete the last element(the least recently used), and insert the new one
	if cache.capacity == len(cache.key2node) {
		elem := cache.data.Back()
		cache.data.Remove(elem)
		delete(cache.key2node, elem.Value.(*Node).key)
	}
	node := &Node{key, value}
	elem := cache.data.PushFront(node)
	cache.key2node[key] = elem
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
}
