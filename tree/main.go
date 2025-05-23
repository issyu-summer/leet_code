package main

import (
	"fmt"
	"math"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//1.Tree Related----------------------------------------------------------------------------------

// Tree-DFS Related------------------------------------------------------------------------------
// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := depth(root.Left)
		r := depth(root.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}
	fmt.Println(depth(root))
	return res
}

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := depth(root.Left)
		r := depth(root.Right)
		if l == -1 || r == -1 {
			return -1
		}
		if math.Abs(float64(l-r)) > 1 {
			return -1
		}
		return max(l, r) + 1
	}
	return depth(root) != -1
}

// 相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
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
	return check(root.Left, root.Right)
}

// 翻转二叉树
func flipTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = flipTree(root.Right), flipTree(root.Left)
	return root
}

// 从前序与中序遍历中构造二叉树
func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//前序：root,[]left,[]right
	//中序：[]left,root,[]right
	var rootIdx int
	for rootIdx < len(inorder) {
		if preorder[0] == inorder[rootIdx] {
			break
		}
		rootIdx++
	}
	leftInOrder := inorder[:rootIdx]
	rightInOrder := inorder[rootIdx+1:]
	leftPreOrder := preorder[1 : len(leftInOrder)+1]
	rightPreOrder := preorder[len(leftInOrder)+1:]
	node := &TreeNode{Val: preorder[0]}
	node.Left = buildTree(leftPreOrder, leftInOrder)
	node.Right = buildTree(rightPreOrder, rightInOrder)
	return node
}

// 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}

// 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}

// 求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	var helper func(root *TreeNode, sum int) int
	helper = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := root.Val + prevSum*10
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return helper(root.Left, sum) + helper(root.Right, sum)
	}
	return helper(root, 0)
}

// 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	var res = math.MinInt
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(maxGain(root.Left), 0)
		r := max(maxGain(root.Right), 0)
		res = max(res, root.Val+l+r)
		return max(l, r) + root.Val
	}
	maxGain(root)
	return res
}

// 二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归展开左右子树
	flatten(root.Left)
	flatten(root.Right)

	// 保存当前右子树
	right := root.Right

	// 将左子树移到右子树位置
	root.Right = root.Left
	root.Left = nil

	// 找到新右子树的末端
	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}

	// 拼接原来的右子树
	curr.Right = right
}

// Tree-BFS Related---------------------------------------------------------------------------
// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		leveSize := len(q)
		var level []int
		for i := 0; i < leveSize; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// 二叉树的锯齿形遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var leftToRight = true
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		var level []int
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if leftToRight {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		leftToRight = !leftToRight
		res = append(res, level)
	}
	return res
}

// 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if i == levelSize {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}

// 2.BST Related----------------------------------------------------------------------------------
// 将有序数组转换为BST
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return root
}

// 3.BackTrack Related----------------------------------------------------------------------------------
// 路径总和I
func hasPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	//此时还未减去root.Val
	if root.Left == nil && root.Right == nil && root.Val == target {
		return true
	}
	return hasPathSum(root.Left, target-root.Val) || hasPathSum(root.Right, target-root.Val)
}

// 路径总和II
func pathSumII(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func(root *TreeNode, path []int, target int)
	backTrack = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		//还未减
		if root.Left == nil && root.Right == nil && root.Val == target {
			res = append(res, append(append([]int{}, path...), root.Val))
			return
		}
		backTrack(root.Left, append(path, root.Val), target-root.Val)
		backTrack(root.Right, append(path, root.Val), target-root.Val)
	}
	backTrack(root, []int{}, target)
	return res
}

// 路径总和III
func pathSumIII(root *TreeNode, target int) int {
	var res int
	prefix := map[int]int{0: 1}
	var backTrack func(root *TreeNode, sum int)
	backTrack = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		res += prefix[sum+root.Val-target]
		prefix[sum+root.Val]++
		backTrack(root.Left, sum+root.Val)
		backTrack(root.Right, sum+root.Val)
		prefix[sum+root.Val]--
	}
	backTrack(root, 0)
	return res
}
