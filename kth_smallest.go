package main

import "math"

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	ans := math.MaxInt
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if k == 0 {
			return
		}
		k--
		ans = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
