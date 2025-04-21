package _024_12_19

func detectCycle(head *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := m[cur]; ok {
			return cur
		} else {
			m[cur] = true
		}
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//需要一点推导
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}
