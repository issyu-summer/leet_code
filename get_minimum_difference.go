package main

import "math"

func main() {

}

func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt, -1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -1 && root.Val-pre < ans {
			ans = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
