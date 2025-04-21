package main

import "fmt"

func main() {
	fmt.Println(greaterOrEqual([]int{1, 3, 3, 5, 7, 7, 9, 11}, 7, true))
	fmt.Println(greaterOrEqual([]int{1, 3, 3, 5, 7, 7, 9, 11}, 7, false))
	fmt.Println(lessOrEqual([]int{1, 3, 3, 5, 6, 8, 9, 11}, 7, false))
	fmt.Println(lessOrEqual([]int{1, 3, 3, 5, 6, 8, 9, 11}, 7, true))
}

func binarySearch(arr []int, target int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		mid := i + (j-i)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

// 第一个大于等于目标值的位置
func lowerBound(arr []int, target int) int {
	i, j := 0, len(arr)
	for i < j {
		mid := i + (j-i)>>1
		if arr[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

// 第一个大于目标值的位置
func upperBound(arr []int, target int) int {
	i, j := 0, len(arr)
	for i < j {
		mid := i + (j-i)>>1
		if arr[mid] <= target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

// 统一左右边界[lower,upper)
// 1, 3, 3, 5, 7, 7, 9, 1
// lower(大于等于) 4
// upper(大于) 6
func greaterOrEqual(arr []int, target int, equal bool) int {
	i, j := 0, len(arr)
	for i < j {
		mid := i + (j-i)>>1
		if arr[mid] < target || (!equal && arr[mid] == target) {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

// 第一个小于（lower-1），要lower:3
// 第一个小于等于（lower包含等于，所以为lower），要upper:4
func lessOrEqual(arr []int, target int, equal bool) int {
	i, j := 0, len(arr)

	for i < j {
		mid := i + (j-i)/2
		if arr[mid] < target {
			i = mid + 1 // mid 可能是解，继续在右半部分查找
		} else {
			j = mid // arr[mid] > target，继续在左半部分查找
		}
	}
	if equal {
		return i
	}
	return i - 1 // 返回第一个小于等于 target 的位置
}
