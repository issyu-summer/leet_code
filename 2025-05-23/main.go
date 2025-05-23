package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next

		//从后向前进行连接
		first.Next = second.Next
		second.Next = first
		pre.Next = second

		pre = first
	}
	return dummy.Next
}
