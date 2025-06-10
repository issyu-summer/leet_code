package main

func main() {

}

type LRUCache struct {
	n int
	L *Node
	R *Node
	m map[int]*Node
}

type Node struct {
	key   int
	val   int
	left  *Node
	right *Node
}

func initNode(k, v int) *Node {
	return &Node{
		key: k,
		val: v,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		n: capacity,
		L: &Node{},
		R: &Node{},
		m: map[int]*Node{},
	}
	lru.L.right = lru.R
	lru.R.left = lru.L
	return lru
}

func (this *LRUCache) Get(key int) int {
	p, ok := this.m[key]
	if !ok {
		return -1
	}

	this.remove(p)
	this.insert(p)
	return p.val
}

func (this *LRUCache) Put(key int, value int) {
	p, ok := this.m[key]
	if ok {
		p.val = value
		this.remove(p)
		this.insert(p)
		return
	}

	if len(this.m) == this.n {
		q := this.R.left
		this.remove(q)
		delete(this.m, q.key)
		// delete(q)
	}
	node := initNode(key, value)
	this.m[key] = node
	this.insert(node)
}

func (this *LRUCache) remove(p *Node) {
	p.right.left = p.left
	p.left.right = p.right
}

func (this *LRUCache) insert(p *Node) {
	p.right = this.L.right
	p.left = this.L
	this.L.right.left = p
	this.L.right = p
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
