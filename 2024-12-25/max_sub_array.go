package _024_12_25

import "math"

func maxSubArray(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	m := l + (r-l)>>1
	a := helper(nums, l, m)
	b := helper(nums, m+1, r)
	c := crossSum(nums, l, m, r)
	return max(a, b, c)
}

func crossSum(nums []int, l, m, r int) int {
	lSum := nums[m]
	sum := 0
	for i := m; i >= l; i-- {
		sum += nums[i]
		if sum > lSum {
			lSum = sum
		}
	}
	rSum := nums[m+1]
	sum = 0
	for i := m + 1; i <= r; i++ {
		sum += nums[i]
		if sum > rSum {
			rSum = sum
		}
	}
	return lSum + rSum
}

func max(list ...int) int {
	ans := math.MinInt64
	for _, v := range list {
		if v > ans {
			ans = v
		}
	}
	return ans
}
