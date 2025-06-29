package main

func main() {

}

type LRUCache struct {
	n     int
	L, R  *Node
	cache map[int]*Node
}

type Node struct {
	key, val    int
	left, right *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		n:     capacity,
		L:     &Node{},
		R:     &Node{},
		cache: map[int]*Node{},
	}
	lru.L.right = lru.R
	lru.R.left = lru.L
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.insertToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key, val int) {
	if node, ok := this.cache[key]; ok {
		node.val = val
		this.remove(node)
		this.insertToHead(node)
		return
	}
	if len(this.cache) >= this.n {
		tail := this.R.left
		this.remove(tail)
		delete(this.cache, tail.key)
	}
	node := &Node{key: key, val: val}
	this.cache[key] = node
	this.insertToHead(node)
}

func (this *LRUCache) remove(node *Node) {
	first, _, third := node.left, node, node.right
	first.right = third
	third.left = first
}

func (this *LRUCache) insertToHead(node *Node) {
	first, second, third := this.L, node, this.L.right
	first.right = second
	second.left = first
	second.right = third
	third.left = second
}
