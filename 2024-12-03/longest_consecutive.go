package _024_12_03

import (
	"sort"
)

// [1,2,3,4] 从小往大
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	result := 0
	for num := range m {
		if !m[num-1] {
			var (
				current    = num
				tempResult = 1
			)
			for m[current+1] {
				current++
				tempResult++
			}
			if result < tempResult {
				result = tempResult
			}
		}
	}
	return result
}

func longestConsecutive2(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	result := 1
	tempResult := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i]-1 == nums[i-1] {
			tempResult++
			if result < tempResult {
				result = tempResult
			}
		} else {
			tempResult = 1
		}
	}
	return result
}
