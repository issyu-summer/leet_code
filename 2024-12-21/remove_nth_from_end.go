package _024_12_21

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	m := map[int]*ListNode{}
	cur := head
	idx := 0
	for cur != nil {
		m[idx] = cur
		idx++
		cur = cur.Next
	}
	if idx == n {
		return head.Next
	}
	//长度是length，倒数第n个索引是length-n
	m[idx-n-1].Next = m[idx-n-1].Next.Next
	return head
}

// 同样适用于数组
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	first, second := head, pre
	for i := 0; i < n; i++ {
		first = first.Next
	}
	//找到倒数第n个的前一个节点
	for first != nil {
		second = second.Next
		first = first.Next
	}
	second.Next = second.Next.Next
	return pre.Next
}
