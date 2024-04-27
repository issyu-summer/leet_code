package main

func main() {

}

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	for k > 0 {
		temp := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = temp
		k--
	}
}
