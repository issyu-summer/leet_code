package main

func main() {

}

func isSymmetric(root *TreeNode) bool {
	return root == nil || recur(root.Left, root.Right)
}

func recur(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	rTree := []*TreeNode{root.Left, root.Right}
	for len(rTree) > 0 {
		lNode, rNode := rTree[0], rTree[1]
		rTree = rTree[2:]
		if lNode == nil && rNode == nil {
			continue
		}
		if lNode == nil || rNode == nil {
			return false
		}
		if lNode.Val != rNode.Val {
			return false
		}
		rTree = append(rTree, lNode.Left, rNode.Right)
		rTree = append(rTree, lNode.Right, rNode.Left)
	}
	return len(rTree) == 0
}
