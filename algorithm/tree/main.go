package main

import (
	"math"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}
	_ = helper(root)
	return res
}

func isBalanced(root *TreeNode) bool {
	var balance = true
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		if math.Abs(float64(left-right)) > 1 {
			balance = false
			return -1
		}
		if left == -1 || right == -1 {
			balance = false
			return -1
		}
		return max(left, right) + 1
	}
	_ = helper(root)
	return balance
}
