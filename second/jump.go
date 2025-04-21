package main

import "math"

func main() {
	jump([]int{2, 3, 1, 1, 4})
}

func jump(nums []int) int {
	return greedy(nums)
}

func greedy(nums []int) int {
	cnt, curEnd, curFarthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		curFarthest = max(curFarthest, i+nums[i])
		if i == curEnd {
			cnt++
			curEnd = curFarthest
			if curEnd >= len(nums)-1 {
				break
			}
		}
	}
	return cnt
}

func dp(nums []int) int {
	f := make([]int, len(nums))
	for i, _ := range f {
		f[i] = math.MaxInt32
	}
	f[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			f[j] = min(f[j], f[i]+1)
		}
	}
	return f[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
