package main

import "container/list"

type Node struct {
	key, value, freq int
}

type LFUCache struct {
	capacity    int
	freq2List   map[int]*list.List
	key2Element map[int]*list.Element
	minFreq     int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:    capacity,
		freq2List:   make(map[int]*list.List),
		key2Element: make(map[int]*list.Element),
		minFreq:     0,
	}
}

func (cache *LFUCache) Get(key int) int {
	elem, ok := cache.key2Element[key]
	if !ok {
		return -1
	}

	node := elem.Value.(*Node)
	oldFreq := node.freq
	newFreq := node.freq + 1
	node.freq = newFreq
	// move the list element from old freq list to new freq list
	cache.freq2List[oldFreq].Remove(elem)
	if cache.freq2List[newFreq] == nil {
		cache.freq2List[newFreq] = list.New()
	}
	newElem := cache.freq2List[newFreq].PushFront(node)
	cache.key2Element[key] = newElem
	if oldFreq == cache.minFreq && cache.freq2List[oldFreq].Len() == 0 {
		cache.minFreq++
	}
	return node.value
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}
	elem, ok := cache.key2Element[key]
	if ok {
		node := elem.Value.(*Node)
		node.value = value
		cache.Get(key)
		return
	}

	if cache.capacity == len(cache.key2Element) {
		// remove the least frequently used element
		// update both freq2List and key2Element
		elem := cache.freq2List[cache.minFreq].Back()
		cache.freq2List[cache.minFreq].Remove(elem)
		delete(cache.key2Element, elem.Value.(*Node).key)
	}

	node := &Node{
		key:   key,
		value: value,
		freq:  1,
	}
	if cache.freq2List[1] == nil {
		cache.freq2List[1] = list.New()
	}
	cache.key2Element[key] = cache.freq2List[1].PushFront(node)
	cache.minFreq = 1
}

func main() {

}
