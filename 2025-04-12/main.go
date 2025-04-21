package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	//在in order中找到root的idx:
	//左子树的inorder：root前[:i]
	//右子树的inorder：root后[i+1:]位右子树的中序变
	//左子树的preorder：root后到长度为左子树的inorder长度[1:len(inorder[:i])+1]
	//右子树的preorder：root->left pre order->right pre order[len(inorder[:i])+1:]
	for i < len(inorder) {
		if inorder[i] == preorder[0] {
			break
		}
		i++
	}
	leftTreeInOrder := inorder[:i]
	rightTreeInOrder := inorder[i+1:]
	leftTreePreOrder := preorder[1 : len(inorder[:i])+1]
	rightTreePreOrder := preorder[len(inorder[:i])+1:]
	root.Left = buildTree(leftTreePreOrder, leftTreeInOrder)
	root.Right = buildTree(rightTreePreOrder, rightTreeInOrder)
	return root
}

func rootSum(root *TreeNode, target int) (res int) {
	if root == nil {
		return 0
	}
	if root.Val == target {
		res++
	}
	res += rootSum(root.Left, target-root.Val)
	res += rootSum(root.Right, target-root.Val)
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func pathSumPrefix(root *TreeNode, targetSum int) int {
	count := 0
	prefixSum := map[int]int{0: 1}
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		count += prefixSum[sum-targetSum]
		prefixSum[sum]++
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		prefixSum[sum]--
	}
	dfs(root, 0)
	return count
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func findPathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var backTrack func(root *TreeNode, target int, path []int)
	backTrack = func(root *TreeNode, target int, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		target -= root.Val
		if root.Left == nil && root.Right == nil && target == 0 {
			res = append(res, append([]int{}, path...))
		}
		backTrack(root.Left, target, path)
		backTrack(root.Right, target, path)
	}
	backTrack(root, targetSum, []int{})
	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	//l,r都不为空，则说明p,q分别在2侧，则返回root
	return root
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(maxGain(root.Left), 0)
		r := max(maxGain(root.Right), 0)
		//res是left+right+root
		sum := root.Val + l + r
		if sum > res {
			res = sum
		}
		//最大贡献是root+左或者root加右
		return root.Val + max(l, r)
	}
	maxGain(root)
	return res
}
