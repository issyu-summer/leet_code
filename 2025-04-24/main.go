package main

import (
	"fmt"
	"sort"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	fmt.Println("search path", root)
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	startPrev := dummy
	for i := 0; i < left-1; i++ {
		startPrev = startPrev.Next
	}
	end := dummy
	for i := 0; i < right; i++ {
		end = end.Next
	}
	start := startPrev.Next
	endNxt := end.Next
	fmt.Println("start", start)
	fmt.Println("end", end)
	startPrev.Next = nil
	end.Next = nil
	_ = reverse(start, nil)
	//end变为头，start变为尾
	startPrev.Next = end
	start.Next = endNxt
	return dummy.Next
}

func reverse(start, end *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := start
	for cur != end {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	m := len(matrix)
	n := len(matrix[0])
	top, bottom := 0, m-1
	left, right := 0, n-1
	for left <= right && top <= bottom {
		//top不动，从left走到right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		//right不动，从top到bottom
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		//bottom不动，从right走到left
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
		}
		bottom--
		//left不动，从bottom走到top
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

func lengthOfLIS(nums []int) int {
	greedyAndSearch := func(nums []int) int {
		var path []int
		for i := 0; i < len(nums); i++ {
			if len(path) == 0 || path[len(path)-1] < nums[i] {
				path = append(path, nums[i])
			} else {
				target := nums[i]
				idx := sort.Search(len(path), func(i int) bool {
					return path[i] >= target
				})
				if idx == len(path) {
					continue
				}
				path[idx] = target
			}
		}
		return len(path)
	}
	//var res int
	//n := len(nums)
	////以i结尾的最长子序列长度
	//f := make([]int, n)
	//for i := 0; i < n; i++ {
	//	f[i] = 1
	//}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < i; j++ {
	//		if nums[i] > nums[j] {
	//			f[i] = max(f[i], f[j]+1)
	//		}
	//	}
	//	res = max(res, f[i])
	//}
	//return res
	return greedyAndSearch(nums)
}
