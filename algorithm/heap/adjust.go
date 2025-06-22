package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 2, 6, 6, 9, 10, 299, 323, 22, 345}
	buildMaxHeap := func(nums []int) {
		for i := len(nums)/2 - 1; i >= 0; i-- {
			down(nums, i)
		}
	}
	sort := func(nums []int) {
		buildMaxHeap(nums)
		for i := len(nums) - 1; i >= 0; i-- {
			nums[0], nums[i] = nums[i], nums[0]
			down(nums[:i], 0)
		}
	}
	sort(nums)
	fmt.Println(nums)
}

// push:append->up(i=len(nums)-1)
func up(nums []int, i int) {
	parent := (i - 1) / 2
	if parent < 0 || nums[parent] >= nums[i] {
		return
	}
	nums[parent], nums[i] = nums[i], nums[parent]
	up(nums, parent)
}

// pop:swap(0,len(nums)-1)->nums[:len(nums)-1]->down(i=0)
func down(nums []int, i int) {
	l, r, largest := 2*i+1, 2*i+2, i
	if l < len(nums) && nums[l] >= nums[largest] {
		largest = l
	}
	if r < len(nums) && nums[r] >= nums[largest] {
		largest = r
	}
	if largest == i {
		return
	}
	nums[i], nums[largest] = nums[largest], nums[i]
	down(nums, largest)
}
