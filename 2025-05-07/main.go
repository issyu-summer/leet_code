package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	for i := 0; i < 16; i++ {
		fmt.Println(i, ":", strconv.FormatInt(int64(i), 2), ":", strconv.FormatInt(int64(i), 16))
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := height(root.Left)
		rightHeight := height(root.Right)
		//当前树不平衡
		if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
			return -1
		}
		//左右子树不平衡
		if leftHeight == -1 || rightHeight == -1 {
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return height(root) >= 1
}

func rotate(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			//0,1=2,0
			matrix[i][j] = matrix[n-1-j][i]
			//2,0=3,2
			matrix[n-1-j][i] = matrix[n-1-i][m-1-j]
			//3,2=1,3
			matrix[n-1-i][m-1-j] = matrix[j][m-1-i]
			//1,3=0,1
			matrix[j][m-1-i] = tmp
		}
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		//row max=7,col max=7
		num := (row-1)*7 + (col - 1)
		if num < 40 {
			return 1 + num%10
		}
	}
}

func isValidBST(root *TreeNode) bool {
	var check func(root *TreeNode, lower, upper int) bool
	check = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		return check(root.Left, lower, root.Val) && check(root.Right, root.Val, upper)
	}
	return check(root, math.MinInt, math.MaxInt)
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}
