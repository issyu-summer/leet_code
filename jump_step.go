package main

func main() {

}

func jump(nums []int) int {
	var (
		target      = len(nums) - 1
		maxPosition = 0
		end         = 0
		step        = 0
	)
	for i := 0; i <= target; i++ {
		if end >= len(nums)-1 {
			return step
		}
		maxPosition = max3(maxPosition, i+nums[i])
		if i == end {
			//若我走过我当前能走的尽头，我一定要向前走一步
			step++
			end = maxPosition
		}
	}
	return step
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
