package main

import (
	"fmt"
	"sort"
)

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	idx := sort.Search(m*n, func(i int) bool {
		row := i / n
		col := i % n
		return matrix[row][col] >= target
	})
	row := idx / n
	col := idx % n
	if row < m && col < n && matrix[row][col] == target {
		fmt.Println("row", "col", row, col, "val", matrix[row][col])
		return true
	}
	return false
}

func partitionLabels(s string) []int {
	var res []int
	lastPos := [26]int{}
	//有多个a，会被最后一个a覆盖掉
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		fmt.Println("idx", i, "char", string(c))
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
			fmt.Println("end update", end)
		}
		if i == end {
			fmt.Println("res added", end-start+1)
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
