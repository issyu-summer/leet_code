package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	//src->dst
	nodeMap := map[*Node]*Node{}
	for cur != nil {
		nodeMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	preHead := &Node{}
	pre := preHead
	for cur != nil {
		pre.Next = nodeMap[cur]
		pre.Next.Random = nodeMap[cur.Random]
		cur = cur.Next
		pre = pre.Next
	}
	return preHead.Next
}
