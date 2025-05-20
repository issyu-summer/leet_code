package main

import (
	"fmt"
)

func main() {
	prefix := prefixSum([]int{1, 2, 3, 4, 5})
	fmt.Println(prefix)
	sum := rangeSum(prefix, 1, 3)
	fmt.Println(sum)
	cnt := subArraySum([]int{1, 2, 3, 4, 5}, 7)
	fmt.Println(cnt)
	fmt.Println(subArraySumWithPrefix([]int{1, 2, 3, 4, 5, 3, 2, 1}, 6))
}

func prefixSum(arr []int) []int {
	prefix := append([]int{}, 0)
	for i := 1; i <= len(arr); i++ {
		prefix = append(prefix, prefix[i-1]+arr[i-1])
	}
	return prefix
}

func rangeSum(prefix []int, l, r int) int {
	return prefix[r+1] - prefix[l]
}

func subArraySum(nums []int, target int) [][]int {
	var res [][]int
	var backTrack func([]int, int, int)
	backTrack = func(path []int, target int, start int) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(nums); i++ {
			backTrack(append(path, nums[i]), target-nums[i], i+1)
		}
	}
	backTrack([]int{}, target, 0)
	return res
}

func subArraySumWithPrefix(nums []int, target int) ([][]int, int) {
	var cnt int
	var res [][]int
	prefix := map[int][]int{
		0: {-1},
	}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if indices, ok := prefix[sum-target]; ok {
			for _, start := range indices {
				res = append(res, nums[start+1:i+1])
				cnt++
			}
		}
		prefix[sum] = append(prefix[sum], i)
	}
	return res, cnt
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func pathSumII(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func(root *TreeNode, path []int, sum int)
	backTrack = func(root *TreeNode, path []int, sum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		sum += root.Val
		if root.Left == nil && root.Right == nil && sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		backTrack(root.Left, path, sum)
		backTrack(root.Right, path, sum)
	}
	backTrack(root, []int{}, 0)
	return res
}

func pathSumIII(root *TreeNode, target int) int {
	var res [][]int
	var backTrack func(*TreeNode, []int, int)
	var prefix = map[int][]int{0: {-1}}
	backTrack = func(root *TreeNode, path []int, sum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		sum += root.Val
		if idxList, ok := prefix[sum-target]; ok {
			for _, start := range idxList {
				res = append(res, append([]int{}, path[start+1:]...))
				fmt.Println(path[start+1:])
			}
		}
		prefix[sum] = append(prefix[sum], len(path)-1)
		backTrack(root.Left, path, sum)
		backTrack(root.Right, path, sum)
		prefix[sum] = prefix[sum][:len(prefix[sum])-1]
	}
	backTrack(root, []int{}, 0)
	return len(res)
}
