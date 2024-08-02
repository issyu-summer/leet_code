package main

import "fmt"

func main() {
	fmt.Print(findMin([]int{1, 2, 3, 4, 5, 6, 7}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
