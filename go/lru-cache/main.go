package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	size     int
	itemList *list.List
	items    map[interface{}]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:     capacity,
		itemList: list.New(),
		items:    make(map[interface{}]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.items[key]; ok {
		this.itemList.MoveToFront(v)

		if v.Value.(*entry) != nil {
			return v.Value.(*entry).value
		}

		return -1
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.items[key]; ok {
		this.itemList.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}

	v := &entry{key, value}
	e := this.itemList.PushFront(v)
	this.items[key] = e

	if this.itemList.Len() > this.size {
		v := this.itemList.Back()

		if v != nil {
			this.itemList.Remove(v)
			k := v.Value.(*entry).key
			delete(this.items, k)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)

	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}
