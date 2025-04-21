package _024_12_23

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{Next: head}
	a := pre
	for a != nil {
		b := a
		for i := 0; i < k; i++ {
			b = b.Next
			if b == nil {
				return pre.Next
			}
		}
		c := b.Next
		b.Next = nil
		h, t := reverse(a.Next)
		a.Next = h
		t.Next = c
		a = t
	}
	return pre.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	//翻转前
	var tail *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = tail
		tail = cur
		cur = nxt
	}
	//翻转后
	return tail, head
}
