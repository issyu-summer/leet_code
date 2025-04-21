package main

func main() {

}

func rob(nums []int) int {
	return robMem(0, nums, make(map[int]int, len(nums)+1))
}

func robMem(i int, nums []int, memo map[int]int) int {
	if i >= len(nums) {
		return 0
	}
	if val, ok := memo[i]; ok {
		return val
	}
	//rob current and next-next
	robI := nums[i] + robMem(i+2, nums, memo)
	//rob next
	skipRobI := robMem(i+1, nums, memo)
	memo[i] = max(robI, skipRobI)
	return memo[i]
}

func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func robing(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	f := make([]int, len(nums))
	f[0] = nums[0]
	f[1] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		f[i] = max(f[i-1], f[i-2]+nums[i])
	}
	return f[len(nums)-1]
}
