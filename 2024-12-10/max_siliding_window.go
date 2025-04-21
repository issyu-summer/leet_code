package _024_12_10

import "math"

// 变长滑动窗口
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	maxNum := math.MinInt
	for i := 0; i < k; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	ans = append(ans, maxNum)
	checkAns := func(num int) bool {
		return num > maxNum
	}
	check := func(l, r int) bool {
		return r-l+1 > k
	}
	for l, r := 0, k; r < len(nums); r++ {
		for ; l <= r && check(l, r); l++ {
		}
		if checkAns(nums[r]) {
			maxNum = nums[r]
			ans = append(ans, maxNum)
		} else {
			ans = append(ans, maxNum)
		}
	}
	return ans
}

// 定长
func maxSlidingWindow1(nums []int, k int) []int {
	ans := make([]int, 0)
	q := []int{}
	for i := 0; i < len(nums); i++ {
		for len(q) != 0 && nums[q[len(q)-1]] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		//idx低于下界面
		if q[0] <= i-k {
			q = q[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}
