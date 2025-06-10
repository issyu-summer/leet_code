package main

type Node struct {
	key   int
	val   int
	left  *Node
	right *Node
}
type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func NewLRUCache(capacity int) LRUCache {
	this := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	this.head.right = this.tail
	this.tail.left = this.head
	return this
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.addToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.remove(node)
		this.addToHead(node)
		return
	}
	if len(this.cache) >= this.capacity {
		back := this.tail.left
		this.remove(back)
		delete(this.cache, back.key)
	}
	node := &Node{key: key, val: value}
	this.addToHead(node)
	this.cache[key] = node
}

func (this *LRUCache) remove(node *Node) {
	first, _, third := node.left, node, node.right
	first.right = third
	third.left = first
}

func (this *LRUCache) addToHead(node *Node) {
	first, second, third := this.head, node, this.head.right

	first.right = second
	second.left = first

	second.right = third
	third.left = second
}
