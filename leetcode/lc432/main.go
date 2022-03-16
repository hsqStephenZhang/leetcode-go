package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	keys map[string]struct{}
	cnt  int
}

type AllOne struct {
	allNodes *list.List
	key2Node map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{
		allNodes: list.New(),
		key2Node: make(map[string]*list.Element),
	}
}

func (t *AllOne) Inc(key string) {
	elem, ok := t.key2Node[key]
	// 增加新的结点，不需要在原先基础上删除
	if !ok {
		// 存在 cnt=1 的结点
		if t.allNodes.Front() != nil && t.allNodes.Front().Value.(*Node).cnt == 1 {
			t.allNodes.Front().Value.(*Node).keys[key] = struct{}{}
			t.key2Node[key] = t.allNodes.Front()
			return
		} else { // 不存在，需要创建，并添加到链表中
			newNode := &Node{
				keys: make(map[string]struct{}),
				cnt:  1,
			}
			newNode.keys[key] = struct{}{}
			t.allNodes.PushFront(newNode)
			t.key2Node[key] = t.allNodes.Front()
		}
	} else { // 需要注意删除操作
		// 不存在 cnt+1 的结点
		if elem.Next() == nil || elem.Next().Value.(*Node).cnt != elem.Value.(*Node).cnt+1 {
			newNode := &Node{
				keys: make(map[string]struct{}),
				cnt:  elem.Value.(*Node).cnt + 1,
			}
			t.allNodes.InsertAfter(newNode, elem)
		}
		delete(elem.Value.(*Node).keys, key)
		elem.Next().Value.(*Node).keys[key] = struct{}{}
		t.key2Node[key] = elem.Next()
		if len(elem.Value.(*Node).keys) == 0 {
			t.allNodes.Remove(elem)
		}
	}
}

func (t *AllOne) Dec(key string) {
	elem := t.key2Node[key]
	node := elem.Value.(*Node)
	if node.cnt == 1 {
		delete(node.keys, key)
		delete(t.key2Node, key)
	} else {
		if elem.Prev() == nil || elem.Prev().Value.(*Node).cnt != elem.Value.(*Node).cnt-1 {
			newNode := &Node{
				keys: make(map[string]struct{}),
				cnt:  elem.Value.(*Node).cnt - 1,
			}
			t.allNodes.InsertBefore(newNode, elem)
		}
		delete(elem.Value.(*Node).keys, key)
		elem.Prev().Value.(*Node).keys[key] = struct{}{}
		t.key2Node[key] = elem.Prev()
	}
	if len(node.keys) == 0 {
		t.allNodes.Remove(elem)
	}
}

func (t *AllOne) GetMinKey() string {
	front := t.allNodes.Front()
	if front == nil {
		return ""
	}
	for k := range front.Value.(*Node).keys {
		return k
	}
	return ""
}

func (t *AllOne) GetMaxKey() string {
	back := t.allNodes.Back()
	if back == nil {
		return ""
	}
	for k := range back.Value.(*Node).keys {
		return k
	}
	return ""
}

func main() {
	allOne := Constructor()
	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("b")
	allOne.Inc("c")
	allOne.Inc("c")
	allOne.Inc("c")
	allOne.Dec("b")
	allOne.Dec("b")
	fmt.Println(allOne.GetMinKey())
	allOne.Dec("a")
	fmt.Println(allOne.GetMaxKey())
	fmt.Println(allOne.GetMinKey())
}
