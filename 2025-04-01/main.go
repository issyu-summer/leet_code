package main

import (
	"fmt"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	fmt.Println(root.Val)
	return root
}

func isSymmetric(root *TreeNode) bool {
	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
}

var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	depth(root, res)
	return res
}

func depth(root *TreeNode, res int) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left, res)
	r := depth(root.Right, res)
	res = max(res, l+r)
	return max(l, r) + 1
}

func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := calculateDepth(root.Left)
	r := calculateDepth(root.Right)
	return max(l, r) + 1
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		level := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		level := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		ans = append(ans, level)
	}
	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans
}

func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(nums []int, left, right int) *TreeNode
	helper = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)>>1
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(nums, left, mid-1)
		root.Right = helper(nums, mid+1, right)
		return root
	}
	return helper(nums, 0, len(nums)-1)
}
