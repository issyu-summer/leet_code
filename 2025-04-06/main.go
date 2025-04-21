package main

import (
	"fmt"
)

func main() {

}

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	var backTrack func([]int)
	backTrack = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
			fmt.Println("add to res：", tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backTrack([]int{})
	return res
}

func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
func climb(n int) int {
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}

func rob(nums []int) int {
	//last
	a := 0
	//cur
	b := nums[0]
	for i := 2; i <= len(nums); i++ {
		a, b = b, max(a+nums[i-1], b)
	}
	return a
}
