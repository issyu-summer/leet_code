package _024_12_29

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		ans = append(ans, root.Val)
		helper(root.Right)
	}
	helper(root)
	return ans
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	//先进后出，左中右=>右左中
	stack := []*TreeNode{root}
	ans := make([]int, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		ans = append(ans, node.Val)
	}
	return ans
}
