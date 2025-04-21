package _024_12_25

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeHelper(lists, 0, len(lists)-1)
}

func mergeHelper(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := l + (r-l)>>1
	l1 := mergeHelper(lists, l, mid)
	l2 := mergeHelper(lists, mid+1, r)
	return mergeTowLists(l1, l2)
}

func mergeTowLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTowLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTowLists(l1, l2.Next)
		return l2
	}
}
