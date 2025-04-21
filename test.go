package main

import "fmt"

// firstLessOrEqual 查找第一个小于等于 target 的位置
func firstLessOrEqual(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1 // mid 可能是解，继续在右半部分查找
		} else {
			right = mid // arr[mid] > target，继续在左半部分查找
		}
	}

	return left // 返回第一个小于等于 target 的位置
}

func main() {
	arr := []int{1, 3, 3, 5, 7, 7, 9, 11}
	target := 7

	result := firstLessOrEqual(arr, target)
	if result >= 0 && result < len(arr) {
		fmt.Printf("第一个小于等于 %d 的元素在数组中的位置是: %d\n", target, result)
	} else {
		fmt.Printf("数组中没有元素小于等于 %d\n", target)
	}
}
