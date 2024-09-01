package main

func main() {

}

func removeDuplicatesII(nums []int) int {
	remove := func(nums []int, k int) int {
		if len(nums) < k {
			return len(nums)
		}
		i, j := k, k
		for j < len(nums) {
			if nums[i-k] != nums[j] {
				nums[i] = nums[j]
				i++
			}
			j++
		}
		return i
	}
	return remove(nums, 2)
}
