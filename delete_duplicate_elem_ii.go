package main

func main() {

}

func removeDuplicatesII(nums []int) int {
	process := func(nums []int, k int) int {
		u := 0
		for _, num := range nums {
			if u < k || nums[u-k] != num {
				nums[u] = num
				u++
			}
		}
		return u
	}
	return process(nums, 2)
}
