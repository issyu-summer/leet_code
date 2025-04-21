package main

import (
	"cmp"
	"sort"
)

func main() {

}

//func permute(nums []int, unique bool) [][]int {
//	if unique {
//		sort.Ints(nums)
//	}
//	res := make([][]int, 0)
//	used := make([]bool, len(nums))
//	var backTrack func([]int)
//	backTrack = func(path []int) {
//		if len(nums) == len(path) {
//			tmp := make([]int, len(nums))
//			copy(tmp, path)
//			res = append(res, tmp)
//			return
//		}
//		for i := 0; i < len(nums); i++ {
//			if used[i] || (unique && i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
//				continue
//			}
//			used[i] = true
//			path = append(path, nums[i])
//			backTrack(path)
//			//回溯
//			path = path[:len(path)-1]
//			used[i] = false
//		}
//	}
//	backTrack([]int{})
//	return res
//}

//func permuteUnique(nums []int) [][]int {
//	return permute(nums, true)
//}

// TreeNode 使用自动回溯，代码更为清晰
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 树型回溯
// 113.路径总和
func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var fc func(root *TreeNode, target int, path []int)
	fc = func(root *TreeNode, target int, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil && target == root.Val {
			res = append(res, append([]int{}, path...))
			return
		}
		fc(root.Left, target-root.Val, path)
		fc(root.Right, target-root.Val, path)
	}
	fc(root, target, []int{})
	return res
}

func slice[T cmp.Ordered](s []T) []T {
	return s
}

// func max[T cmp.Ordered](x T, y ...T) T
// 排列和组合需要分开
func backTrackFc[T cmp.Ordered](checkFc func(path []T) (bool, bool), arr []T, combination bool, do, undo func(path []T, idx int)) [][]T {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var res [][]T
	var fc func(path []T, start int, used uint64)
	fc = func(path []T, start int, used uint64) {
		retDirectly, resultReturn := checkFc(path)
		if resultReturn {
			res = append(res, append([]T{}, path...))
			return
		}
		if retDirectly {
			return
		}
		if combination {
			for i := start; i < len(arr); i++ {
				mask := uint64(1 << i)
				if mask&used != 0 || i > 0 && arr[i] == arr[i-1] && used&uint64(1<<(i-1)) == 0 {
					continue
				}
				do(path, i)
				fc(append(path, arr[i]), i+1, used|mask)
				undo(path, i)
			}
		} else {
			for i := 0; i < len(arr); i++ {
				mask := uint64(1 << i)
				if mask&used != 0 || i > 0 && arr[i] == arr[i-1] && used&uint64(1<<(i-1)) == 0 {
					continue
				}
				fc(append(path, arr[i]), start, used|mask)
			}
		}
	}
	fc([]T{}, 0, 0)
	return res
}

// 数组型回溯
// 46.全排列，无重复
func permute(nums []int) [][]int {
	check := func(nums []int) func(path []int) (bool, bool) {
		return func(path []int) (bool, bool) {
			return false, len(path) == len(nums)
		}
	}
	res := backTrackFc(check(nums), nums, false)
	return res
	//var res [][]int
	//var fc func([]int, uint64)
	//fc = func(path []int, used uint64) {
	//	if len(path) == len(nums) {
	//		res = append(res, append([]int{}, path...))
	//		return
	//	}
	//	for i := 0; i < len(nums); i++ {
	//		mask := uint64(1 << i)
	//		if used&mask != 0 {
	//			continue
	//		}
	//		fc(append(path, nums[i]), used|mask)
	//	}
	//}
	//fc([]int{}, 0)
	//return res
}

// 全排列II，有重复
func permuteWithRepeatNum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var fc func([]int, uint64)
	fc = func(path []int, used uint64) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			mask := uint64(1 << i)
			if used&mask != 0 || (i > 0 && nums[i] == nums[i-1] && used&(1<<(i-1)) == 0) {
				continue
			}
			fc(append(path, nums[i]), used|mask)
		}
	}
	fc([]int{}, 0)
	return res
}

// 子集
func subsets(nums []int) [][]int {
	var res [][]int
	var fc func([]int, int)
	fc = func(path []int, start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			fc(append(path, nums[i]), i+1)
		}
	}
	fc([]int{}, 0)
	return res
}

// 子集，含重复元素
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var fc func(path []int, start int, used uint64)
	fc = func(path []int, start int, used uint64) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			mask := uint64(1 << i)
			if i > 0 && nums[i] == nums[i-1] && used&uint64(1<<(i-1)) == 0 {
				continue
			}
			fc(append(path, nums[i]), i+1, used|mask)
		}
	}
	fc([]int{}, 0, 0)
	return res
}

// 组合总数，无重复，但可重复
func combinationSum(candidates []int, target int) [][]int {
	check := func(target int) func([]int) (bool, bool) {
		return func(path []int) (bool, bool) {
			return target < 0, target == 0
		}
	}
	do := func(target int) func(path []int, idx int) {
		return func(path []int, idx int) {
			target -= candidates[idx]
		}
	}
	undo := func(target int) func(path []int, idx int) {
		return func(path []int, idx int) {
			target += candidates[idx]
		}
	}
	res := backTrackFc(check(target), candidates, true, do(target), undo(target))
	//var res [][]int
	//var fc func([]int, int, int)
	//fc = func(path []int, start int, target int) {
	//	if target == 0 {
	//		res = append(res, append([]int{}, path...))
	//		return
	//	}
	//	if target < 0 {
	//		return
	//	}
	//	for i := start; i < len(candidates); i++ {
	//		fc(append(path, candidates[i]), i, target-candidates[i])
	//	}
	//}
	//fc([]int{}, 0, target)
	return res
}

// 组合总数2，有重复，不可重复选
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var fc func(path []int, start int, target int, used uint64)
	fc = func(path []int, start int, target int, used uint64) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			mask := uint64(1 << i)
			//[]int{1, 1, 2, 2, 3}
			//前一个选中，当前没有选中，组合为1,3
			//前一个没有选中，当前一个选中，组合为1,3
			//前一个选中，当前一个选中，组合为1,1,3
			//两种组合会重复
			if i > 0 && candidates[i] == candidates[i-1] && used&uint64(1<<(i-1)) == 0 {
				continue
			}
			fc(append(path, candidates[i]), i+1, target-candidates[i], used|mask)
		}
	}
	fc([]int{}, 0, target, 0)
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
	fc = func(bytes []byte, open int, close int) {
		if len(bytes) == n*2 {
			res = append(res, string(bytes))
		}
		if open < n {
			fc(append(bytes, '('), open+1, close)
		}
		if close < open {
			fc(append(bytes, ')'), open, close+1)
		}
	}
	fc([]byte{}, 0, 0)
	return res
}
