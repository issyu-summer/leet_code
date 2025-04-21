package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	cnt := map[*ListNode]struct{}{}
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := cnt[cur]; ok {
			return true
		}
		cnt[cur] = struct{}{}
	}
	return false
}

func floyd(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
