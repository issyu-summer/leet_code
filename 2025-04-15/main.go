package main

import (
	"fmt"
)

func main() {

}

func generateParenthesis(n int) []string {
	var res []string
	var backTrack func(string, int, int, int)
	backTrack = func(current string, open, close, max int) {
		if len(current) == max*2 {
			fmt.Println("add res", current)
			res = append(res, current)
			return
		}
		if open < max {
			fmt.Println("open<max current", current+"(")
			backTrack(current+"(", open+1, close, max)
		}
		if close < open {
			fmt.Println("close<open current", current+")")
			backTrack(current+")", open, close+1, max)
		}
	}
	backTrack("", 0, 0, n)
	return res
}

func generate(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var res []string
	for i := 0; i < n; i++ {
		for _, l := range generate(i) {
			for _, r := range generate(n - 1 - i) {
				res = append(res, "("+l+")"+r)
			}
		}
	}
	return res
}

func exist(board [][]byte, word string) bool {
	var dfs func(board [][]byte, word string, i, j, index int, visited [][]bool) bool
	dfs = func(board [][]byte, word string, i, j, index int, visited [][]bool) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] || board[i][j] != word[index] {
			return false
		}
		visited[i][j] = true
		if dfs(board, word, i+1, j, index+1, visited) ||
			dfs(board, word, i-1, j, index+1, visited) ||
			dfs(board, word, i, j+1, index+1, visited) ||
			dfs(board, word, i, j-1, index+1, visited) {
			return true
		}
		visited[i][j] = false
		return false
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

func existII(board [][]byte, word string) bool {
	var dfs func(board [][]byte, word string, i, j, index int) bool
	dfs = func(board [][]byte, word string, i, j, index int) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[index] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '#'
		if dfs(board, word, i+1, j, index+1) ||
			dfs(board, word, i-1, j, index+1) ||
			dfs(board, word, i, j+1, index+1) ||
			dfs(board, word, i, j-1, index+1) {
			return true
		}
		board[i][j] = tmp
		return false
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
