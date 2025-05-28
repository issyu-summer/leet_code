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

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return nil
	}
	var backTrack func(root *TreeNode, path []int, target int)
	backTrack = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		fmt.Println("path", path, "target", target, "val", root.Val)
		if root.Left == nil && root.Right == nil && target == root.Val {
			res = append(res, append([]int{}, append(path, root.Val)...))
			return
		}
		backTrack(root.Left, append(path, root.Val), target-root.Val)
		backTrack(root.Right, append(path, root.Val), target-root.Val)
	}
	backTrack(root, []int{}, targetSum)
	return res
}

// 1,2,1
func findPeakElement(nums []int) int {
	return sort.Search(len(nums)-1, func(i int) bool {
		return nums[i] > nums[i+1]
	})
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//var lcp func(a, b string) string
	//lcp = func(a, b string) string {
	//	if len(a) > len(b) {
	//		return lcp(b, a)
	//	}
	//	var idx int
	//	for idx < len(a) {
	//		if a[idx] == b[idx] {
	//			idx++
	//		} else {
	//			break
	//		}
	//	}
	//	return a[:idx]
	//}
	var partition func(strs []string, start, end int) string
	partition = func(strs []string, start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := start + (end-start)/2
		left := partition(strs, start, mid)
		right := partition(strs, mid+1, end)
		minL := min(len(left), len(right))
		for i := 0; i < minL; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minL]
	}
	return partition(strs, 0, len(strs)-1)
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	var res int
	for num, _ := range m {
		if !m[num-1] {
			//cur:1
			cur := num
			curLength := 1
			//cur+1:2
			for m[cur+1] {
				//1->2
				cur++
				curLength++
			}
			res = max(res, curLength)
		}
	}
	return res
}
