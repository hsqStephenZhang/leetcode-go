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
	capacity int
	key2node map[int]*list.Element
	data     list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		key2node: make(map[int]*list.Element),
		data:     list.List{},
	}
}

func (cache *LRUCache) Get(key int) int {
	if elem, ok := cache.key2node[key]; ok {
		cache.data.MoveToFront(elem)
		return elem.Value.(*Node).value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if elem, ok := cache.key2node[key]; ok {
		cache.data.MoveToFront(elem)
		node := elem.Value.(*Node)
		node.value = value
		return
	}

	if len(cache.key2node) == cache.capacity {
		last := cache.data.Back()
		cache.data.Remove(last)
		delete(cache.key2node, last.Value.(*Node).key)
	}
	newElem := cache.data.PushFront(&Node{key, value})
	cache.key2node[key] = newElem
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 0)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
