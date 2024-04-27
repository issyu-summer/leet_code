package main

import (
	"fmt"
)

func main() {
	ar := []int{0}
	merge(ar, 0, []int{1}, 1)
	fmt.Print(ar)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, tail := m-1, n-1, m+n-1
	for {
		if i == -1 {
			nums1[tail] = nums2[j]
			j--
		} else if j == -1 {
			nums1[tail] = nums1[i]
			i--
		} else if nums2[j] >= nums1[i] {
			nums1[tail] = nums2[j]
			j--
		} else {
			nums1[tail] = nums1[i]
			i--
		}
		tail--
		if i < 0 && j < 0 {
			return
		}
	}
}
