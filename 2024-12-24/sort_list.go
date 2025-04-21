package _024_12_24

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := middleNode1(head)
	nxt := mid.Next
	mid.Next = nil
	left := sortList(head)
	right := sortList(nxt)
	return mergeTwoSortedList(left, right)
}

// 如果是偶数个，则返回中间两个节点的第1个节点
func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 如果是偶数个，则返回中间两个节点的第2个节点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoSortedList(l1, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}
