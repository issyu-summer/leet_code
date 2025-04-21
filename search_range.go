package main

func main() {
	searchRange([]int{1, 1, 1}, 1)

}

func searchRange(nums []int, target int) []int {
	l := binarySearch(nums, target, true)
	if l < len(nums) && nums[l] == target {
		if l < len(nums)-1 && nums[l+1] == target {
			tmp := l
			for tmp < len(nums)-1 && nums[tmp] == target {
				tmp++
			}
			if tmp == len(nums)-1 && nums[tmp] == target {
				return []int{l, tmp}
			}
			return []int{l, tmp - 1}
		} else if l > 0 && nums[l-1] == target {
			tmp := l
			for tmp < len(nums)-1 && nums[tmp] == target {
				tmp--
			}
			if tmp == 0 && nums[tmp] == target {
				return []int{tmp, l}
			}
			return []int{tmp + 1, l}
		} else {
			return []int{l, l}
		}
	}
	return []int{-1, -1}
}
