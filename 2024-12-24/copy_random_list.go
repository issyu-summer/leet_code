package _024_12_24

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{}
	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := m[node]; ok {
			return n
		}
		newNode := &Node{Val: node.Val}
		m[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}
	return deepCopy(head)
}

// 替身文学
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	//生成替身
	cur := head
	for cur != nil {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
		cur = cur.Next.Next
	}
	//完美替身
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	//独立替身
	newHead := head.Next
	cur = head
	for cur != nil {
		nxt := cur.Next
		cur.Next = cur.Next.Next
		if nxt.Next != nil {
			nxt.Next = nxt.Next.Next
		}
		cur = cur.Next
	}
	return newHead
}
