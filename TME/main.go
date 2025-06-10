package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		fmt.Printf("%d,", node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	fmt.Println()
	q = []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		node.Left, node.Right = node.Right, node.Left
		//fmt.Println(node, node.Left, node.Right)
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	q = []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		fmt.Printf("%d,", node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
}

func lengthOfLIS(nums []int) int {
	var lis []int
	for i := 0; i < len(nums); i++ {
		if len(lis) == 0 || (len(lis) > 0 && lis[len(lis)-1] < nums[i]) {
			lis = append(lis, nums[i])
		} else {
			target := nums[i]
			idx := sort.Search(len(lis), func(i int) bool {
				return lis[i] >= target
			})
			if idx < len(nums) {
				lis[idx] = target
			}
		}
	}
	return len(lis)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	midNxt := slow.Next
	mid.Next = nil
	var pre, cur *ListNode = nil, midNxt
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	newHead := pre
	a, b := head, newHead
	for b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

func rand7() int {
	return rand.IntN(7) + 1
}

func rand10() int {
	for {
		row, col := rand7(), rand7()
		idx := (row-1)*7 + col
		if idx <= 40 {
			return (idx-1)%10 + 1
		}
	}
}

// 数组中的重复数据、原地hash
func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	fmt.Println(nums)
	var res []int
	for k := 0; k < len(nums); k++ {
		if nums[k] != k+1 {
			res = append(res, nums[k])
		}
	}
	return res
}

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	f := make([][]int, len(text1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
			}
		}
	}
	return f[len(text1)][len(text2)]
}

// TreeNode 二叉树的右视
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}

// 编辑距离
func minDistance(word1 string, word2 string) int {
	f := make([][]int, len(word1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(f); i++ {
		f[i][0] = i
	}
	for j := 0; j < len(f[0]); j++ {
		f[0][j] = j
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1], f[i][j]) + 1
			}
		}
	}
	return f[len(word1)+1][len(word2)+1]
}

// 缺失的第一个正数
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return -1
}

// 合并区间
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || (len(res) > 0 && res[len(res)-1][1] < intervals[i][0]) {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

// 最大正方形
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
				res = max(res, f[i+1][j+1])
			}
		}
	}
	return res * res
}

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
	mid := slow
	midNxt := slow.Next
	mid.Next = nil
	l1 := sortList(head)
	l2 := sortList(midNxt)
	return mergeList(l1, l2)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil || root.Left.Val != root.Right.Val {
		return false
	}
	return isSymmetric(root.Left) && isSymmetric(root.Right)
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	res := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if res < x && res*res == x {
		return res
	}
	return res - 1
}

func longestValidParentheses(s string) int {
	var stack []int
	var res int
	for r := 0; r < len(s); r++ {
		if s[r] == '(' {
			stack = append(stack, r)
		} else {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				var curLen int
				if len(stack) == 0 {
					curLen = r - 0 + 1
				} else {
					curLen = r - stack[len(stack)-1]
				}
				res = max(res, curLen)
			} else {
				stack = append(stack, r)
			}
		}
	}
	return res
}

func compareVersion(version1 string, version2 string) int {
	nums1 := strings.Split(version1, ".")
	nums2 := strings.Split(version2, ".")
	var i, j int
	for i < len(nums1) || j < len(nums2) {
		var a, b int
		if i < len(nums1) {
			a, _ = strconv.Atoi(nums1[i])
		}
		if j < len(nums2) {
			b, _ = strconv.Atoi(nums2[j])
		}
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		i++
		j++
	}
	return 0
}
