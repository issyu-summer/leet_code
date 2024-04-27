package main

func removeDuplicates(nums []int) int {
	n, i, j := len(nums), 1, 1
	if n == 0 {
		return 0
	}
	for j < n {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
