package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	isValidBST2(root)
}

// error
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

// 中序遍历
func isValidBST1(root *TreeNode) bool {
	var dfs func(root *TreeNode, lower, upper int) bool
	dfs = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		return dfs(root.Left, lower, root.Val) && dfs(root.Right, root.Val, upper)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

func isValidBST2(root *TreeNode) bool {
	tmp := math.MinInt
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !dfs(root.Left) {
			return false
		}
		if root.Val <= tmp {
			return false
		}
		tmp = root.Val
		if !dfs(root.Right) {
			return false
		}
		return true
	}
	return dfs(root)
}
