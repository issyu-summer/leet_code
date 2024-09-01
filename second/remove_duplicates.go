package main

func main() {

}

func removeDuplicates(nums []int) int {
	//快慢双指针
	i, j := 1, 1
	for j < len(nums) {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
