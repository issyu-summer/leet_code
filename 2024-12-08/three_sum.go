package _024_12_08

import "sort"

// -4,-1,-1,0,1,2
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	for i, _ := range result {
		result[i] = make([]int, 0)
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
			}
		}
	}
	return result
}
