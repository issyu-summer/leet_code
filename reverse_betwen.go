package main

func main() {

}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head
	firstHalfHead, rightNode := preHead, preHead

	for i := 0; i < left-1; i++ {
		firstHalfHead = firstHalfHead.Next
	}
	for i := 0; i < right; i++ {
		rightNode = rightNode.Next
	}

	leftNode := firstHalfHead.Next
	firstHalfHead.Next = nil

	lastHalfHead := rightNode.Next
	rightNode.Next = nil

	//left_node->a->b->...->right_node
	_ = reverseList1(leftNode)

	firstHalfHead.Next = rightNode
	leftNode.Next = lastHalfHead
	return preHead.Next
}

func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
