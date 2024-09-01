package main

func main() {

}

// is this two pointers???
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
