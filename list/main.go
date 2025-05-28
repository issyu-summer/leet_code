package main

import (
	"fmt"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加（链表版本）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	sum, carry := 0, 0
	for l1 != nil || l2 != nil {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		sum %= 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// 删除链表的倒数第N个基点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 合并两个有序链表
func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

// 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var partition func(lists []*ListNode, l, r int) *ListNode
	partition = func(lists []*ListNode, l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		mid := l + (r-l)>>1
		l1 := partition(lists, l, mid)
		l2 := partition(lists, mid+1, r)
		return mergeHelper(l1, l2)
	}
	return partition(lists, 0, len(lists)-1)
}

func mergeHelper(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeHelper(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeHelper(l2.Next, l1)
		return l2
	}
}

// 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next

		first.Next = second.Next
		second.Next = first
		pre.Next = second

		pre = first
	}
	return dummy.Next
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	return reverseHelper(head, nil)
}

// 反转链表II
func reverseBetween(head *ListNode, l, r int) *ListNode {
	dummy := &ListNode{Next: head}
	lHeadPre, lHead, rTail, rTailNxt := dummy, dummy, dummy, dummy
	for i := 0; i < l-1; i++ {
		lHeadPre = lHeadPre.Next
	}
	lHead = lHeadPre.Next
	for i := 0; i < r; i++ {
		rTail = rTail.Next
	}
	rTailNxt = rTail.Next
	fmt.Println(lHeadPre, lHead, rTail, rTailNxt)
	lHeadPre.Next = nil
	rTail.Next = nil
	_ = reverseHelper(lHead, nil)
	lHeadPre.Next = rTail
	fmt.Println(rTailNxt)
	lHead.Next = rTailNxt
	return dummy.Next
}

// 1->2->3->nil，head 1 end nil
// 3->2->1->nil，pre为3,实际是end前面那个
func reverseHelper(head, end *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != end {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// k个一组反转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	end := head
	for i := 0; i < k; i++ {
		if end == nil {
			return head
		}
		end = end.Next
	}
	newHead := reverseHelper(head, end)
	head.Next = reverseKGroup(end, k)
	return newHead
}

// 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			val := fast.Val
			for fast.Next != nil && fast.Next.Val == val {
				fast = fast.Next
			}
			slow.Next = fast
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return dummy.Next
}

// 删除链表中的重复元素II
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
	for fast != nil {
		if fast.Next != nil && fast.Next.Val == fast.Val {
			val := fast.Val
			for fast != nil && fast.Val == val {
				fast = fast.Next
			}
			slow.Next = fast
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return dummy.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 随机链表的复制
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := map[*Node]*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		newNode := &Node{Val: cur.Val}
		nodeMap[cur] = newNode
	}
	for cur := head; cur != nil; cur = cur.Next {
		newNode := nodeMap[cur]
		if cur.Next != nil {
			newNode.Next = nodeMap[cur.Next]
		}
		if cur.Random != nil {
			newNode.Random = nodeMap[cur.Random]
		}
	}
	return nodeMap[head]
}

// 环形链表
func hasCycle(head *ListNode) bool {
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

// 环形链表II
func detectCycle(head *ListNode) *ListNode {
	var exist bool
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			exist = true
			break
		}
	}
	if !exist {
		return nil
	}
	res := head
	for res != slow {
		res = res.Next
		slow = slow.Next
	}
	return res
}

// 重排链表
func reorderList(head *ListNode) {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	midNxt := slow.Next
	slow.Next = nil
	list := reverseHelper(midNxt, nil)

	first, second := head, list
	for second != nil {
		a, b := first.Next, second.Next
		first.Next = second
		second.Next = a
		first, second = a, b
	}
}

// 排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(mid)
	return mergeHelper(l1, l2)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre != nil {
		cur := pre.Next
		if cur != nil && cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}
