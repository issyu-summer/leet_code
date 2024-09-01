package main

func main() {

}

func minSubArrayLen(target int, nums []int) int {
	var (
		ans        = len(nums) + 1
		sum        = 0
		start, end = 0, 0
	)
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			if ans > end-start+1 {
				ans = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
