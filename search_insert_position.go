package main

func main() {
	searchInsert([]int{1, 3, 5}, 3)

}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
