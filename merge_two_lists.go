package main

func main() {

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead
	h1 := list1
	h2 := list2
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}
	if h1 != nil {
		cur.Next = h1
	}
	if h2 != nil {
		cur.Next = h2
	}
	return preHead.Next
}
