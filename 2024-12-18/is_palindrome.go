package _024_12_18

import "fmt"

func isPalindrome(head *ListNode) bool {
	ints := make([]int, 0)
	cur := head
	for cur != nil {
		ints = append(ints, cur.Val)
		cur = cur.Next
	}
	l, r := 0, len(ints)-1
	for l < r {
		if ints[l] != ints[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	m := middleNode(head)
	rHead := reverseList(m)
	a := head
	b := rHead
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

// slow走1步，fast走2步，fast走到头，则slow走了一半
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	fmt.Println(cur == nil)
	return pre
}
