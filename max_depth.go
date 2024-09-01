package main

func main() {

}

func maxDepth(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return maxDepth(root.Left) + maxDepth(root.Right) + 1
}
