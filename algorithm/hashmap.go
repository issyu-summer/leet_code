package main

import (
	"fmt"
	"hash/fnv"
)

// 键值对节点
type Entry struct {
	Key, Val string
	Next     *Entry
}

type HashMap struct {
	buckets    []*Entry
	capacity   int
	size       int
	loadFactor float64
	threshold  int
}

func Constructor(capacity int, loadFactor float64) *HashMap {
	if capacity == 0 {
		capacity = 16
	}
	if loadFactor == 0 {
		loadFactor = 0.75
	}
	return &HashMap{
		buckets:    make([]*Entry, capacity),
		capacity:   capacity,
		loadFactor: loadFactor,
		threshold:  int(float64(capacity) * loadFactor),
	}
}

func (hm *HashMap) bucketIdx(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % hm.capacity
}

func (hm *HashMap) Put(key, val string) {
	if hm.size >= hm.threshold {
		hm.resize()
	}
	idx := hm.bucketIdx(key)
	entry := &Entry{key, val, nil}
	if hm.buckets[idx] == nil {
		hm.buckets[idx] = entry
		hm.size++
		return
	}
	cur := hm.buckets[idx]
	for {
		if cur.Key == key {
			cur.Val = val
			return
		}
		if cur.Next == nil {
			cur.Next = entry
			break
		}
		cur = cur.Next
	}
}

func (hm *HashMap) Get(key string) (string, bool) {
	idx := hm.bucketIdx(key)
	cur := hm.buckets[idx]
	for cur != nil {
		if cur.Key == key {
			return cur.Val, true
		}
		cur = cur.Next
	}
	return "", false
}

func (hm *HashMap) Delete(key string) {
	idx := hm.bucketIdx(key)
	cur := hm.buckets[idx]
	var pre *Entry
	for cur != nil {
		if cur.Key == key {
			if pre == nil {
				hm.buckets[idx] = cur.Next
			} else {
				pre.Next = cur.Next
			}
			hm.size--
			return
		}
		pre = cur
		cur = cur.Next
	}
}

func (hm *HashMap) resize() {
	hm.capacity *= 2
	newBuckets := make([]*Entry, hm.capacity)
	for _, head := range hm.buckets {
		cur := head
		for cur != nil {
			entry := &Entry{cur.Key, cur.Val, nil}
			newIdx := hm.bucketIdx(cur.Key)
			if newBuckets[newIdx] == nil {
				newBuckets[newIdx] = entry
			} else {
				entry.Next = newBuckets[newIdx]
				newBuckets[newIdx] = entry
			}
			cur = cur.Next
		}
	}
	hm.buckets = newBuckets
	hm.threshold = int(float64(hm.capacity) * hm.loadFactor)
}

func main() {
	hm := Constructor(8, 0.75)
	hm.Put("apple", "1")
	fmt.Println(hm.Get("apple"))
	hm.Put("apple", "2")
	fmt.Println(hm.Get("apple"))
	hm.Put("banana", "3")
	fmt.Println(hm.Get("banana"))
	for i := 0; i < 100; i++ {
		hm.Put(fmt.Sprint(i), fmt.Sprint(i))
	}
	fmt.Println(hm.capacity, hm.threshold, len(hm.buckets))
}
