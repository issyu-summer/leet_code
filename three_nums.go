package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum > 0 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return result
}
