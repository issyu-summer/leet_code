package _024_12_15

func maxSubArray(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		if ans < f[i] {
			ans = f[i]
		}
	}
	return ans
}

func maxSubArray1(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}
