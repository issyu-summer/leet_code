package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 34, 5, 5, 3434, 65664, 2323, 4, 60, 99, 98, 97, 90, 9999}
	quickSort(nums, 0, len(nums)-1, len(nums)-5)
	fmt.Println(nums[len(nums)-5:])
}

func partition(nums []int, l, r int) int {
	//随机化pivot
	pivot := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func quickSort(nums []int, l, r, k int) {
	pivotIdx := partition(nums, l, r)
	if pivotIdx == k {
		return
	} else if pivotIdx < k {
		quickSort(nums, pivotIdx+1, r, k)
	} else {
		quickSort(nums, l, pivotIdx-1, k)
	}
}
