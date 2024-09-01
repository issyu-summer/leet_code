package main

import "math"

func main() {
	minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	})
}

func minimumTotal(triangle [][]int) int {
	ans := make([][]int, len(triangle))
	ans[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i]); j++ {
			ans[i][j] = ans[i-1][j]min(ans[i][j], triangle[i][j])
		}
		ans[i] = ans[i-1] + minVal(triangle[i][i], triangle[i][i-1])
	}
	return ans[len(triangle)-1]
}

func minVal(arr ...int) int {
	ans := math.MaxInt64
	for _, v := range arr {
		if v < ans {
			ans = v
		}
	}
	return ans
}
