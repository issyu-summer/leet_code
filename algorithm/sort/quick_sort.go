package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	nums := []int{1, 4, 3, 3, 2, 5, 6, 5, 5, 5}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func partition(nums []int, l, r int) int {
	//随机化pivot
	randIdx := l + rand.IntN(r-l+1)
	nums[randIdx], nums[r] = nums[randIdx], nums[r]

	pivot := nums[r]
	i := l
	//最后交换r
	for j := l; j < r; j++ {
		//改为>则是逆序
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivotIdx := partition(nums, l, r)
	quickSort(nums, l, pivotIdx-1)
	quickSort(nums, pivotIdx+1, r)
}
