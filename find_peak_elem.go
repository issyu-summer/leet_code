package main

import "math"

func main() {
	findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})

}

func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i == len(nums) || i == -1 {
			return math.MinInt64
		}
		return nums[i]
	}
	l, r := 0, len(nums)-1
	for {
		mid := l + (r-l)>>1
		if get(mid) > get(mid+1) && get(mid) > get(mid-1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}
