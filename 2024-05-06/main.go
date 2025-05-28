package main

import (
	"bytes"
	"sort"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, prevSum int) int
	dfs = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := root.Val + prevSum*10
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	dfs(root, 0)
	return dfs(root, 0)
}

func isSymmetric(root *TreeNode) bool {
	var check func(left *TreeNode, right *TreeNode) bool
	check = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func searchRange(nums []int, target int) []int {
	start := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if start == len(nums) || (start < len(nums) && nums[start] != target) {
		return []int{-1, -1}
	}
	end := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})
	return []int{start, end - 1}
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backTrack func(path []int, start, target int)
	backTrack = func(path []int, start, target int) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			backTrack(append(path, candidates[i]), i, target-candidates[i])
		}
	}
	backTrack([]int{}, 0, target)
	return res
}

func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			//括号中的内容出栈
			var str []byte
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				str = append([]byte{stack[len(stack)-1]}, str...)
				stack = stack[:len(stack)-1]
			}
			//[出栈
			stack = stack[:len(stack)-1]
			//计数
			var sum int
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				sum = int(stack[len(stack)-1]-'0') + sum*10
				stack = stack[:len(stack)-1]
			}
			//重复
			stack = append(stack, bytes.Repeat(str, sum)...)
		}
	}
	return string(stack)
}
