package main

import "math"

func main() {
	search([]int{4, 5, 6, 7, 0, 1, 2}, 1)
}

func search(nums []int, target int) int {
	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if get(mid) == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target >= get(0) && target < get(mid) {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if nums[mid] <= nums[len(nums)-1] {
			if target > get(mid) && target <= get(len(nums)-1) {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
