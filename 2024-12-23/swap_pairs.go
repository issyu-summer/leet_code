package _024_12_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{Next: head}
	a := pre
	for a.Next != nil && a.Next.Next != nil {
		b := a.Next
		c := b.Next
		a.Next = c
		b.Next = c.Next
		c.Next = b
		a = b
	}
	return pre.Next
}
