package main

func main() {

}

func lengthOfLIS(nums []int) int {
	ans := 0
	// 1.init status
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	// 2.status transfer
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func lengthOfLISII(nums []int) int {
	ans := 0
	n := len(nums)
	tails := make([]int, n)
	for k := 0; k < n; k++ {
		i, j := 0, k
		for i < j {
			mid := (i + j) >> 1
			if tails[mid] < nums[k] {
				i = mid + 1
			} else {
				j = mid
			}
		}
		tails[i] = nums[k]
		if ans == j {
			ans++
		}
	}
	return ans
}
