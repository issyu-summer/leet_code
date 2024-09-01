package main

func main() {

}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func (b *BSTNode) Insert(val int) {
	if val < b.Val {
		if b.Left == nil {
			b.Left = &BSTNode{Val: val}
		} else {
			b.Left.Insert(val)
		}
	} else {
		if b.Right == nil {
			b.Right = &BSTNode{Val: val}
		} else {
			b.Right.Insert(val)
		}
	}
}

func (b *BSTNode) Search(val int) *BSTNode {
	if b == nil || b.Val == val {
		return b
	}
	if val < b.Val {
		return b.Left.Search(val)
	}
	return b.Right.Search(val)
}

func (b *BSTNode) FindMin() *BSTNode {
	current := b
	if current.Left != nil {
		current = current.Left
	}
	return current
}

func (b *BSTNode) Delete(val int) *BSTNode {
	if b == nil {
		return nil
	}
	if val < b.Val {
		b.Left = b.Left.Delete(val)
	} else if val > b.Val {
		b.Right = b.Right.Delete(val)
	} else {
		//有一个子节点，直接用子节点代替
		if b.Left == nil {
			return b.Right
		} else if b.Right == nil {
			return b.Left
		}

		//有两个子节点，交互要删除的值与其中继后续（右子树的最小值）
		minNode := b.Right.FindMin()
		b.Val = minNode.Val
		b.Right = b.Right.Delete(minNode.Val)
	}
	return b
}
