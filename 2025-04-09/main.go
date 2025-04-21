package main

import (
	"fmt"
)

func main() {
	var res int
	if res == 0 {
		res++
	}
	if res == 1 {
		fmt.Println(res)
	}

	var tmp int
	if tmp == 0 {
		tmp++
	} else if tmp == 1 {
		fmt.Println(tmp)
	}
}

func singleNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func majorityElement(nums []int) int {
	res := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
		}
		if nums[i] == res {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	//short,long
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[m][n]
}

func longestCommonSubsequenceII(text1 string, text2 string) int {
	//short,long
	memo := map[[2]int]int{}
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}
		if val, ok := memo[[2]int{i, j}]; ok {
			return val
		}
		if text1[i] == text2[j] {
			memo[[2]int{i, j}] = f(i+1, j+1) + 1
		} else {
			memo[[2]int{i, j}] = max(f(i+1, j), f(i, j+1))
		}
		return memo[[2]int{i, j}]
	}
	return f(0, 0)
}
