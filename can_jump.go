package main

func main() {

}

func canJump(nums []int) bool {
	m := 0
	for i := 0; i < len(nums); i++ {
		if m < i {
			return false
		}
		m = max1(m, i+nums[i])
	}
	return true
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
