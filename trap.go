package main

func main() {
	trapIII([]int{4, 2, 0, 3, 2, 5})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func trap(height []int) int {
	var (
		ans               = 0
		leftMax, rightMax = make([]int, len(height)), make([]int, len(height))
	)
	leftMax[0], rightMax[len(rightMax)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i := 0; i < len(height); i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func trapII(height []int) int {
	var (
		i, j              = 0, len(height) - 1
		leftMax, rightMax = 0, 0
		ans               = 0
	)
	for i < j {
		if height[i] < height[j] {
			//左边是短板
			if height[i] > leftMax {
				leftMax = height[i]
			} else {
				ans += leftMax - height[i]
			}
			i++
		} else {
			//右边是短板
			if height[j] > rightMax {
				rightMax = height[j]
			} else {
				ans += rightMax - height[j]
			}
			j--
		}
	}
	return ans
}

func trapIII(height []int) int {
	ans := 0
	stack := []int{0}
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			ans += (i - left - 1) * (min(height[left], height[i]) - height[top])
		}
		stack = append(stack, i)
	}
	return ans
}
