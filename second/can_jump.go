package main

func main() {

}

func canJump(nums []int) bool {
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if i > f[i-1] {
			return false
		}
		f[i] = max(f[i-1], i+nums[i])
	}
	return f[len(nums)-1] >= len(nums)-1
}
