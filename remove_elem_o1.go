package main

import "fmt"

func main() {
	ar := []int{1}
	e := removeElement(ar, 1)
	fmt.Print(e, ar)
}

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)
	// len-1需要边界条件需要包含等号，要多判断一次
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j]
			j--
		} else {
			i++
		}
	}
	return i
}
