package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
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
	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	return dummy.Next
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)>>1
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwoList(l1, l2)
}

func reorderList(head *ListNode) {
	middle := func(head *ListNode) *ListNode {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}
	reverse := func(start *ListNode, end *ListNode) *ListNode {
		var pre, cur *ListNode = nil, start
		for cur != end {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		return pre
	}
	mid := middle(head)
	midNxt := mid.Next
	mid.Next = nil
	pre := reverse(midNxt, nil)
	first, second := head, pre
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}

}

func addStrings(num1 string, num2 string) string {
	var res []byte
	i, j := len(num1)-1, len(num2)-1
	var carry uint8 = 0
	for i >= 0 || j >= 0 || carry > 0 {
		var a, b uint8
		if i >= 0 {
			a = num1[i] - '0'
			i--
		}
		if j >= 0 {
			b = num2[j] - '0'
			j--
		}
		sum := a + b + carry
		carry = sum / 10
		sum = sum % 10
		res = append([]byte{sum + '0'}, res...)
	}
	return string(res)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		a = a.Next
		b = b.Next
		if a == nil && b == nil {
			return nil
		}
		if b == nil {
			b = headA
		}
		if a == nil {
			a = headB
		}
	}
	return a
}

func mergeInterval(intervals [][]int) [][]int {
	var res [][]int
	//按照start排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		//当前start比end小，则没有重叠
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		}
	}
	return res
}

func trap(height []int) int {
	var res int
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	//边界
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]
	//loop left/right
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i := 0; i < len(height); i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

func minDistance(word1 string, word2 string) int {
	//前i个字符转换为前j个字符所需要的最小操作次数
	f := make([][]int, len(word1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(word2)+1)
	}
	//全部删除
	for i := 0; i < len(word1)+1; i++ {
		f[i][0] = i
	}
	//全部添加
	for i := 0; i < len(word2)+1; i++ {
		f[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			}
		}
	}
	return f[len(word1)][len(word2)]
}
