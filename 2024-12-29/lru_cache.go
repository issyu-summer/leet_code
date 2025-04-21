package _024_12_29

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	//map用来快速查找
	cache map[int]*list.Element
	//缓存访问顺序
	list *list.List
}

type entry struct {
	key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entry).Value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if elem, ok := l.cache[key]; ok {
		elem.Value.(*entry).Value = value
		l.list.MoveToFront(elem)
		return
	}
	if l.list.Len() >= l.capacity {
		oldest := l.list.Back()
		if oldest != nil {
			l.list.Remove(oldest)
			delete(l.cache, oldest.Value.(*entry).key)
		}
	}
	elem := l.list.PushFront(&entry{key, value})
	l.cache[key] = elem
}
