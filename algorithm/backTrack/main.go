package main

import (
	"sort"
)

func main() {

}

// 全排列I
func permuteI(nums []int, unique bool) [][]int {
	res := make([][]int, 0)
	var backTrack func(path []int, used uint64)
	backTrack = func(path []int, used uint64) {
		if len(nums) == len(path) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			masking := uint64(1 << i)
			if used&masking != 0 {
				continue
			}
			backTrack(append(path, nums[i]), used|masking)
		}
	}
	backTrack([]int{}, 0)
	return res
}

// 全排列II，有重复
func permuteII(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var backTrack func(path []int, used uint64)
	backTrack = func(path []int, used uint64) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			masking := uint64(1 << i)
			if used&masking != 0 || (i > 0 && nums[i] == nums[i-1] && used&uint64(1<<(i-1)) == 0) {
				continue
			}
			backTrack(append(path, nums[i]), used|masking)
		}
	}
	backTrack([]int{}, 0)
	return res
}

// 子集
func subsets(nums []int) [][]int {
	var res [][]int
	var backTrack func(path []int, start int)
	backTrack = func(path []int, start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			backTrack(append(path, nums[i]), i+1)
		}
	}
	backTrack([]int{}, 0)
	return res
}

// 子集，含重复元素
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var backTrack func(path []int, start int, used uint64)
	backTrack = func(path []int, start int, used uint64) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			masking := uint64(1 << i)
			if i > 0 && nums[i] == nums[i-1] && used&uint64(1<<(i-1)) == 0 {
				continue
			}
			backTrack(append(path, nums[i]), i+1, used|masking)
		}
	}
	backTrack([]int{}, 0, 0)
	return res
}

// 组合总数，无重复，但可重复
func combinationSumI(candidates []int, target int) [][]int {
	var res [][]int
	var backTrack func(path []int, start int, target int)
	backTrack = func(path []int, start int, target int) {
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

// 组合总数2，有重复，不可重复选
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var backTrack func(path []int, start int, target int, used uint64)
	backTrack = func(path []int, start int, target int, used uint64) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			masking := uint64(1 << i)
			if i > 0 && candidates[i] == candidates[i-1] && used&uint64(1<<(i-1)) == 0 {
				continue
			}
			backTrack(append(path, candidates[i]), i+1, target-candidates[i], used|masking)
		}
	}
	backTrack([]int{}, 0, target, 0)
	return res
}

// 电话号码的组合
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var res []string
	var fc func([]byte, int)
	fc = func(bytes []byte, start int) {
		if start == len(digits) {
			res = append(res, string(bytes))
			return
		}
		chars := m[digits[start]]
		for i := 0; i < len(chars); i++ {
			fc(append(bytes, chars[i]), start+1)
		}
	}
	fc([]byte{}, 0)
	return res
}

// 括号生成
func generateParenthesis(n int) []string {
	var res []string
	var fc func([]byte, int, int)
	fc = func(path []byte, open int, close int) {
		if len(path) == n*2 {
			res = append(res, string(path))
		}
		if open < n {
			fc(append(path, '('), open+1, close)
		}
		if close < open {
			fc(append(path, ')'), open, close+1)
		}
	}
	fc([]byte{}, 0, 0)
	return res
}

// TreeNode 使用自动回溯，代码更为清晰
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func(root *TreeNode, path []int, target int)
	backTrack = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && root.Val == target {
			res = append(res, append([]int{}, append(path, root.Val)...))
		}
		backTrack(root.Left, append(path, root.Val), target-root.Val)
		backTrack(root.Right, append(path, root.Val), target-root.Val)
	}
	backTrack(root, []int{}, target)
	return res
}
