package main

import (
	"fmt"
	"sort"
)

func main() {

}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func numsIsLands(grid [][]byte) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func findJudge(n int, trust [][]int) int {
	in := make([]int, n)
	out := make([]int, n)
	for _, t := range trust {
		out[t[0]-1]++
		in[t[1]-1]++
	}
	for i := 0; i < n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i + 1
		}
	}
	return -1
}

// 笛卡尔积？为何不是i+1，而是start+1
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
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
	var backTrack func(start int, combination string)
	backTrack = func(start int, combination string) {
		if len(digits) == start {
			fmt.Println("res", combination)
			res = append(res, combination)
			return
		}
		digit := digits[start]
		chars := m[digit]
		for i := 0; i < len(chars); i++ {
			backTrack(start+1, combination+string(chars[i]))
		}
	}
	backTrack(0, "")
	return res
}

// 组合总数，可重复
func combinationSum(candidates []int, target int) [][]int {
	//全局变量或者带着一起loop
	if len(candidates) == 0 {
		return [][]int{}
	}
	var (
		res [][]int
	)
	var backTrack func(start int, path []int, sum int)
	backTrack = func(start int, path []int, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			//可重复，每个数字可以用多次
			backTrack(i, path, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backTrack(0, []int{}, 0)
	return res
}

// 组合总数，不可重复
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	//全局变量或者带着一起loop
	if len(candidates) == 0 {
		return [][]int{}
	}
	var (
		res [][]int
	)
	fmt.Println(candidates)
	var backTrack func(start int, path []int, sum int)
	backTrack = func(start int, path []int, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			fmt.Println("res", temp)
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			fmt.Println("idx", i, "val", candidates[i])
			if candidates[i] > target {
				break
			}
			//[1 1 2 5 6 7 10]
			// 1,1,6不会有2组
			// 1,2,5会有2组
			if i > start && candidates[i] == candidates[i-1] {
				fmt.Println("idx-1", i-1, "val_1", candidates[i-1], "idx", i, "val_2", candidates[i])
				continue
			}
			path = append(path, candidates[i])
			//不可重复,每个数字只用一次
			backTrack(i+1, path, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backTrack(0, []int{}, 0)
	return res
}

// 全排列
func permute(nums []int) [][]int {
	var (
		res [][]int
	)
	var backTrack func([]int, []bool)
	backTrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			fmt.Println("add to res", temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTrack(path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backTrack([]int{}, make([]bool, len(nums)))
	return res
}

// 子集合，不可重复
func subsets(nums []int) [][]int {
	var res [][]int
	var backTrack func(int, []int)
	backTrack = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backTrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backTrack(0, []int{})
	return res
}
