package main

import (
	"math"
)

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(grid [][]int, i, j int) int
	dfs = func(grid [][]int, i, j int) int {
		//越界或者不是岛屿
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)
	}
	var res int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(grid, i, j))
			}
		}
	}
	return res
}

func numIslands(grid [][]byte) int {
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
	var res int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res = math.MinInt
	type Node struct {
		TreeNode *TreeNode
		Idx      int
	}
	q := []*Node{{root, 1}}
	for len(q) > 0 {
		levelSize := len(q)
		leftIdx := q[0].Idx
		rightIdx := q[levelSize-1].Idx
		res = max(res, rightIdx-leftIdx)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if node.TreeNode.Left != nil {
				q = append(q, &Node{
					TreeNode: node.TreeNode.Left,
					Idx:      node.Idx * 2,
				})
			}
			if node.TreeNode.Right != nil {
				q = append(q, &Node{
					TreeNode: node.TreeNode.Right,
					Idx:      node.Idx*2 + 1,
				})
			}
		}
	}
	return res
}

func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		f[i][0] = 1
	}
	for j := 0; j < n; j++ {
		f[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, maxProd, minProd := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}
		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])
		res = max(res, maxProd)
	}
	return res
}

func maxProductF(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxF, minF := make([]int, len(nums)), make([]int, len(nums))
	maxF[0] = nums[0]
	minF[0] = nums[0]
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxF[i-1], minF[i-1] = minF[i-1], maxF[i-1]
		}
		maxF[i] = max(nums[i], maxF[i-1]*nums[i])
		minF[i] = min(nums[i], minF[i-1]*nums[i])
		res = max(res, maxF[i])
	}
	return res
}

func hasPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && target == root.Val {
		return true
	}
	return hasPathSum(root.Left, target-root.Val) || hasPathSum(root.Right, target-root.Val)
}

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func(root *TreeNode, path []int, target int)
	backTrack = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && target == root.Val {
			res = append(res, append(append([]int{}, path...), root.Val))
			return
		}
		backTrack(root.Left, append(path, root.Val), target-root.Val)
		backTrack(root.Right, append(path, root.Val), target-root.Val)
	}
	backTrack(root, []int{}, target)
	return res
}

func pathSumIII(root *TreeNode, target int) int {
	var res [][]int
	if root == nil {
		return len(res)
	}
	var backTrack func(root *TreeNode, path []int, target int)
	backTrack = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		if target == root.Val {
			res = append(res, append(append([]int{}, path...), root.Val))
		}
		backTrack(root.Left, append(path, root.Val), target-root.Val)
		backTrack(root.Right, append(path, root.Val), target-root.Val)
	}

	var preOrder func(root *TreeNode) int
	preOrder = func(root *TreeNode) int {
		res = [][]int{}
		if root == nil {
			return 0
		}
		backTrack(root, []int{}, target)
		return len(res) + preOrder(root.Left) + preOrder(root.Right)
	}
	return preOrder(root)
}

func pathSumIIIOptimize(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	var backTrack func(root *TreeNode, target int) int
	backTrack = func(root *TreeNode, target int) int {
		if root == nil {
			return 0
		}
		cnt := 0
		if target == root.Val {
			cnt = 1
		}
		return cnt + backTrack(root.Left, target-root.Val) + backTrack(root.Right, target-root.Val)
	}

	var preOrder func(root *TreeNode) int
	preOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return backTrack(root, target) + preOrder(root.Left) + preOrder(root.Right)
	}
	return preOrder(root)
}

func pathSumIIIWithPath(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func(root *TreeNode, path []int, target int)
	backTrack = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		if target == root.Val {
			res = append(res, append(append([]int{}, path...), root.Val))
		}
		backTrack(root.Left, append(path, root.Val), target-root.Val)
		backTrack(root.Right, append(path, root.Val), target-root.Val)
	}
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		backTrack(root, []int{}, target)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return res
}

func pathSumIIIWithPathOptimize(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func(*TreeNode, []int, int)
	var prefix = map[int][]int{0: {-1}}
	backTrack = func(root *TreeNode, path []int, sum int) {
		if root == nil {
			return
		}
		if idxList, ok := prefix[sum+root.Val-target]; ok {
			for _, start := range idxList {
				res = append(res, append(append([]int{}, path[start+1:]...), root.Val))
			}
		}
		prefix[sum+root.Val] = append(prefix[sum+root.Val], len(path)-1)
		backTrack(root.Left, append(path, root.Val), sum+root.Val)
		backTrack(root.Right, append(path, root.Val), sum+root.Val)
		prefix[sum+root.Val] = prefix[sum+root.Val][:len(prefix[sum+root.Val])-1]
	}
	backTrack(root, []int{}, 0)
	return res
}

func pathSumIIIWithPathOptimizeCnt(root *TreeNode, target int) int {
	var res int
	var backTrack func(*TreeNode, int)
	var prefix = map[int]int{0: 1}
	backTrack = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		if cnt, ok := prefix[sum+root.Val-target]; ok {
			res += cnt
		}
		prefix[sum+root.Val]++
		backTrack(root.Left, sum+root.Val)
		backTrack(root.Right, sum+root.Val)
		prefix[sum+root.Val]--
	}
	backTrack(root, 0)
	return res
}
