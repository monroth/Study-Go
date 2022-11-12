package main

import (
	"container/list"
)

type mapPair struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	contains map[int]*list.Element
	order    list.List
}

func Constructor(capacity int) LRUCache {
	var l list.List
	out := LRUCache{capacity, make(map[int]*list.Element), l}
	return out
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.contains[key]; ok {
		this.order.MoveToFront(node)
		return node.Value.(mapPair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, newValue int) {
	if node, ok := this.contains[key]; ok {
		this.order.MoveToFront(node)
		node.Value = mapPair{key, newValue}
		return
	}
	if this.order.Len() == this.capacity {
		back := this.order.Back()
		delete(this.contains, back.Value.(mapPair).key)
		this.order.Remove(back)
	}
	this.order.PushFront(mapPair{key, newValue})
	this.contains[key] = this.order.Front()

}

func main() {
}
