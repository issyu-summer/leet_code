package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	lMap := map[*ListNode]bool{}
	for head != nil {
		if _, ok := lMap[head]; ok {
			return true
		}
		lMap[head] = true
		head = head.Next
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
