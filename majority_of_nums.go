package main

func main() {
	majorityElement([]int{3, 3, 4, 4, 4, 4, 5, 5})
}
func majorityElement(nums []int) int {
	majority, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
		}
		if majority == nums[i] {
			count++
		} else {
			count--
		}
	}
	return majority
}
