package main

func canJump2(nums []int) bool {
	last := len(nums) - 1
	step := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= step {
			last = i
			step = 1
		} else {
			step++
		}
	}
	return last == 0
}
