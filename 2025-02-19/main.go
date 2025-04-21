package _025_02_19

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func checkSymmetricTree(root *TreeNode) bool {
	var check func(l, r *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil || l.Val != r.Val {
			return false
		}
		return check(l.Left, r.Right) && check(l.Right, r.Left)
	}
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func flipTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = flipTree(root.Right), flipTree(root.Left)
	return root
}

// binary search tree
func getMinimumDifference(root *TreeNode) int {
	//全局变量或者指针
	type ans struct {
		ans int
		pre int
	}
	var dfs func(ans *ans, root *TreeNode)
	dfs = func(ans *ans, root *TreeNode) {
		if root == nil {
			return
		}
		dfs(ans, root.Left)
		if ans.pre != -1 && root.Val-ans.pre < ans.ans {
			ans.ans = root.Val - ans.pre
		}
		ans.pre = root.Val
		dfs(ans, root.Right)
	}
	a := &ans{math.MaxInt, -1}
	dfs(a, root)
	return a.ans
}

func merge(nums1, m int, nums2 []int, n int) []int {
	i, j := 0, 0
	for i < m && j < n {

	}
}
