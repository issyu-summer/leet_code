package main

import (
	"fmt"
)

func main() {

}

func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
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
			//f[j]+=f[j-1]
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	//写几个装填，有助于写出方程
	//f[0][0] = grid[0][0]
	//f[0][1] = grid[0][0] + grid[0][1]
	//f[1][0] = grid[0][0] + grid[1][0]
	//f[1][1] = min(f[1][0], f[0][1]) + grid[1][1]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[i][j] = grid[i][j]
			} else if i == 0 {
				f[i][j] = f[i][j-1] + grid[i][j]
			} else if j == 0 {
				f[i][j] = f[i-1][j] + grid[i][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
			}
			fmt.Println(i, j, f[i][j])
		}
	}
	return f[m-1][n-1]
}

func longestPalindrome(s string) string {
	expand := func(s string, i, j int) (int, int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
		}
		return i + 1, j - 1
	}
	var end, start int
	for i := 0; i < len(s); i++ {
		i1, j1 := expand(s, i, i)
		i2, j2 := expand(s, i, i+1)
		if j1-i1 > end-start {
			start, end = i1, j1
		}
		if j2-i2 > end-start {
			start, end = i2, j2
		}
	}
	return s[start : end+1]
}

func longestPalindromeII(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		f[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			f[i][i+1] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
			fmt.Println(i, j, f[i][j])
		}
	}

	res := ""
	for i := n - 3; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if f[i][j] && len(res) < j-i+1 {
				res = s[i : j+1]
			}
		}
	}
	return res
}
