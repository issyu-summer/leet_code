package main

import (
	"fmt"
)

func main() {

}

func setZeros(matrix [][]int) {
	m := map[[2]int]bool{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				m[[2]int{i, j}] = true
			}
		}
	}
	for idx, isZero := range m {
		if isZero {
			fmt.Println(idx)
			//set column zero,use i
			for i := 0; i < len(matrix); i++ {
				matrix[i][idx[0]] = 0
			}
			//set row zero,use j
			for j := 0; j < len(matrix[0]); j++ {
				matrix[idx[1]][j] = 0
			}
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i, j := m-1, 0; i >= 0 && j <= n-1; {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

func rotate(matrix [][]int, target int) {
	m := len(matrix)
	for i := 0; i < m/2; i++ {
		for j := 0; j < (m+1)/2; j++ {
			tmp := matrix[i][j]
			//0,1->0,2
			matrix[i][j] = matrix[m-1-j][i]
			matrix[m-1-j][i] = matrix[m-1-i][m-1-j]
			matrix[m-1-i][m-1-j] = matrix[j][m-1-i]
			matrix[j][m-1-i] = tmp
		}
	}
}

func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	var res []int
	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if t <= b {
			for i := r; i >= l; i-- {
				res = append(res, matrix[b][i])
			}
			b--
		}
		if l <= r {
			for i := b; i >= t; i-- {
				res = append(res, matrix[i][l])
			}
			l++
		}
	}
	return res
}
