package _024_12_19

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	m := map[*ListNode]bool{}
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := m[cur]; !ok {
			m[cur] = true
		} else {
			return true
		}
	}
	return false
}
