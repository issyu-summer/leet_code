package main

func main() {

}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	rTree := []*TreeNode{root}
	for len(rTree) > 0 {
		node := rTree[0]
		rTree = rTree[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			rTree = append(rTree, node.Left)
		}
		if node.Right != nil {
			rTree = append(rTree, node.Right)
		}
	}
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	rTree := []*TreeNode{root}
	for len(rTree) > 0 {
		node := rTree[len(rTree)-1]
		rTree = rTree[:len(rTree)-1]
		if node.Left != nil {
			rTree = append(rTree, node.Left)
		}
		if node.Right != nil {
			rTree = append(rTree, node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}
