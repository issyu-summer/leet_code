package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 6, 7, 8, 9}
	_ = SearchGT(nums, 5)
	//fmt.Println("gt", gt)
	_ = SearchGE(nums, 5)
	//fmt.Println("ge", ge)
	//eq := SearchEQ(nums, 5)
	//fmt.Println("eq", eq)
	//lt := SearchLT(nums, 5)
	//fmt.Println("lt", lt)
	//le := SearchLE(nums, 5)
	//fmt.Println("le", le)
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target { // 小于
			l = mid + 1
		} else if nums[mid] > target { // 大于
			r = mid - 1
		} else { // 等于
			return mid
		}
	}
	return l
}

func SearchGT(ar []int, target int) int {
	idx, ok := slices.BinarySearchFunc(ar, target, func(cur, target int) int {
		if cur > target {
			return 1
		} else {
			return -1
		}
	})
	fmt.Println("gt", idx, "ok?", ok)
	//if !true,i=mid+1
	idx = sort.Search(len(ar), func(i int) bool { return ar[i] > target })
	if idx == len(ar) {
		return -1
	}
	return idx
}

// 大于的都是返回l
func binarySearchGT(nums []int, target int) int { // 大于
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func SearchGE(ar []int, target int) int {
	idx, ok := slices.BinarySearchFunc(ar, target, func(cur, target int) int {
		if cur >= target {
			return 1
		} else {
			return -1
		}
	})
	fmt.Println("ge", idx, "ok?", ok)
	idx = sort.Search(len(ar), func(i int) bool { return ar[i] >= target })
	if idx == len(ar) {
		return -1
	}
	return idx
}

func binarySearchGE(nums []int, target int) int { // 大于等于
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func SearchEQ(a []int, target int) int {
	i := sort.Search(len(a), func(i int) bool { return a[i] >= target })
	if i < len(a) && a[i] == target {
		return i
	}
	return -1
}

func SearchLE(a []int, target int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > target }) - 1
}

// 小于的都是返回r
func binarySearchLE(nums []int, target int) int { // 小于等于
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func SearchLT(a []int, target int) int {
	//if !true,l=mid+1
	return sort.Search(len(a), func(i int) bool { return a[i] >= target }) - 1
}

func binarySearchLT(nums []int, target int) int { // 小于
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
