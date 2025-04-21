package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现
// root->left->right
func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}

// left->root->right
func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	println(root.Val)
	inOrderTraversal(root.Right)
}

// left-right-root
func postOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraversal(root.Left)
	postOrderTraversal(root.Right)
	println(root.Val)
}

// 栈实现
func preOrderTraversal2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		//弹出栈顶元素并访问
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		println(node.Val)
		//栈是先进后出，所以right先入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func bfs1(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Println(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println()
	}
}
