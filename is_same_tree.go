package main

func main() {

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	pTree := []*TreeNode{p}
	qTree := []*TreeNode{q}
	for len(pTree) > 0 && len(qTree) > 0 {
		pLevelLength := len(pTree)
		qLevelLength := len(qTree)
		if pLevelLength != qLevelLength {
			return false
		}
		for i := 0; i < pLevelLength; i++ {
			pNode := pTree[0]
			qNode := qTree[0]
			if pNode.Left != nil && qNode.Left == nil {
				return false
			}
			if pNode.Right != nil && qNode.Right == nil {
				return false
			}
			if pNode.Left == nil && qNode.Left != nil {
				return false
			}
			if pNode.Right == nil && qNode.Right != nil {
				return false
			}
			if pNode.Val != qNode.Val {
				return false
			}

			pTree = pTree[1:]
			qTree = qTree[1:]
			if pNode.Left != nil {
				pTree = append(pTree, pNode.Left)
			}
			if pNode.Right != nil {
				pTree = append(pTree, pNode.Right)
			}
			if qNode.Left != nil {
				qTree = append(qTree, qNode.Left)
			}
			if qNode.Right != nil {
				qTree = append(qTree, qNode.Right)
			}
		}
	}
	return len(pTree) == 0 && len(qTree) == 0
}
