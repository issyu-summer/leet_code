package main

import (
	"math"
	"sort"
)

func main() {

}

func lengthOfLIS(nums []int) int {
	judge := func(nums []int) int {
		res := make([]int, 0)
		for i := 0; i < len(nums); i++ {
			if len(res) == 0 || res[len(res)-1] < nums[i] {
				res = append(res, nums[i])
			} else {
				target := nums[i]
				idx := sort.Search(len(res), func(i int) bool {
					return res[i] >= target
				})
				if idx == len(res) {
					continue
				}
				res[idx] = nums[i]
			}
		}
		return len(res)
	}
	//n := len(nums)
	//res := make([]int, n)
	//for i := 0; i < n; i++ {
	//	res[i] = 1
	//}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < i; j++ {
	//		计算出以i结尾的最长子序列的长度
	//if nums[i] > nums[j] {
	//	res[i] = max(res[i], res[j]+1)
	//}
	//}
	//}
	//return res[n-1]
	return judge(nums)
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := math.MinInt32
	n := len(nums)
	maxF := make([]int, n)
	minF := make([]int, n)
	maxF[0] = nums[0]
	minF[0] = nums[0]
	for i := 1; i < n; i++ {
		num := nums[i]
		maxF[i] = max(nums[i], max(maxF[i-1]*num, minF[i-1]*num))
		minF[i] = min(nums[i], min(minF[i-1]*num, maxF[i-1]*num))
		if maxF[i] > res {
			res = maxF[i]
		}
	}
	return res
}

func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	f := make([]bool, target+1)
	f[0] = true
	for i := 0; i < n; i++ {
		if nums[i] > target {
			continue
		}
		for j := target; j >= nums[i]; j-- {
			f[j] = f[j] || f[j-nums[i]]
		}
	}
	return f[target]
}

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	res := 0
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				curLen := 0
				if len(stack) == 0 {
					curLen = i + 1
				} else {
					curLen = i - stack[len(stack)-1]
				}
				res = max(res, curLen)
			} else {
				stack = append(stack, i)
			}
		}
	}
	return res
}
