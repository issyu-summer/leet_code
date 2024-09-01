package main

func main() {

}

// 双指针
func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			nums[i] = nums[j]
			j--
		} else {
			i++
		}
	}
	return i
}

func CollisionPointersTemplate(nums []int) int {
	return 0
}

func do() {

}

func logic() bool {
	return true
}

func FastSlowPointersTemplate(nums []int) int {
	i, j := 0, 0
	for j < len(nums) {
		if logic() {
			do()
			i++
		}
		j++
	}
	return i
}
