package _024_12_05

import (
	"math"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	result := math.MinInt
	for i < j {
		area := (j - i) * compare(height[i], height[j], false)
		result = compare(result, area, true)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func compare(i, j int, gte bool) int {
	if gte {
		if i > j {
			return i
		}
		return j
	} else {
		if i < j {
			return i
		}
		return j
	}
}
