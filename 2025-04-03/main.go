package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var check func(root *TreeNode, lower, upper int) bool
	check = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if !(lower < root.Val && root.Val < upper) {
			return false
		}
		return check(root.Left, lower, root.Val) && check(root.Right, root.Val, upper)
	}
	return check(root, math.MinInt, math.MaxInt)
}

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	var inOrderTraverse func(root *TreeNode)
	inOrderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTraverse(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		inOrderTraverse(root.Right)
	}
	inOrderTraverse(root)
	return res
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if i == levelSize-1 {
				res = append(res, cur.Val)
			}
		}
	}
	return res
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	pre := &TreeNode{}
	cur := pre
	var preOrderTraversal func(root *TreeNode)
	preOrderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		cur.Right = &TreeNode{Val: root.Val}
		cur = cur.Right
		preOrderTraversal(root.Left)
		preOrderTraversal(root.Right)
	}
	preOrderTraversal(root)
	curRoot := root
	cur = pre.Right
	for cur != nil {
		fmt.Println("cur", cur.Val)
		curRoot.Val = cur.Val
		if cur.Right != nil {
			curRoot.Left = nil
			if curRoot.Right == nil {
				curRoot.Right = &TreeNode{}
			}
		}
		fmt.Println("curRoot", curRoot.Val)
		cur = cur.Right
		curRoot = curRoot.Right
	}
}

// 小于等于
func searchInsert(nums []int, target int) int {
	search := func(n int, f func(idx int) bool) int {
		i, j := 0, n
		for i < j {
			mid := (i + j) / 2
			if !f(mid) {
				i = mid + 1
			} else {
				j = mid
			}
		}
		return i
	}
	return search(len(nums), func(idx int) bool {
		return target <= nums[idx]
	})
}

func search(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		mid := (i + j) / 2
		if !f(mid) {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	start := sort.Search(n, func(i int) bool {
		fmt.Println("start idx", i, "num", nums[i])
		return nums[i] >= target
	})
	end := sort.Search(n, func(i int) bool {
		fmt.Println("end idx", i, "num", nums[i])
		return nums[i] > target
	}) - 1
	if n == 0 || start == n || nums[start-1] != target {
		return []int{-1, -1}
	}
	return []int{start, end}
}

func searchRotateArr(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 使用一次sort.Search完成查找
	idx := sort.Search(n, func(i int) bool {
		// 当nums[i]和target在同一部分时，正常比较
		if (nums[i] >= nums[0]) == (target >= nums[0]) {
			return nums[i] >= target
		}

		// 当nums[i]在左部分，target在右部分
		return !(nums[i] >= nums[0])
	})

	// 检查找到的索引是否有效
	if idx < n && nums[idx] == target {
		return idx
	}
	return -1
}
