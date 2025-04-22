package main

import (
	"fmt"
)

func main() {
	fmt.Println(10169 + 13612)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//codetop *4

// 1。合并两个有序链表
func mergeTwoListsLoop(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
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
	return pre.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// 2.最长回文子串
func longestPalindrome(s string) string {
	res := ""
	expand := func(s string, i, j int) (int, int) {
		for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
			i--
			j++
		}
		return i + 1, j - 1
	}
	for i := 1; i < len(s); i++ {
		start, end := expand(s, i, i)
		fmt.Println(start, end)
		if len(res) < end-start+1 {
			res = s[start : end+1]
		}
		start, end = expand(s, i, i+1)
		fmt.Println(start, end)
		if len(res) < end-start+1 {
			res = s[start : end+1]
		}
	}
	return res
}

func longestPalindromeDP(s string) string {
	n := len(s)
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		f[i][i] = true
	}
	//handle 2 char
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			f[i][i+1] = true
		}
	}
	//handle >= 3 char
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	res := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if f[i][j] && len(res) < j-i+1 {
				res = s[i : j+1]
			}
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 3.二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		var tmp []int
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

// 4.两束之和
func towSUm(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if lastIdx, ok := m[target-nums[i]]; ok {
			return []int{lastIdx, i}
		}
		m[nums[i]] = i
	}
	return nil
}
