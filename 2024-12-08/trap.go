package _024_12_08

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
