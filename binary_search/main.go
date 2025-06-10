package main

import (
	"sort"
)

func main() {

}

// 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	total := m + n
	half := (total + 1) / 2 //奇数，前2，后1、、、偶数前2，后2
	if m == 0 {
		if half%2 == 1 {
			return float64(nums2[half-1])
		}
		return float64(nums2[half-1]+nums2[half]) / 2.0
	}
	i := sort.Search(m, func(i int) bool {
		//计算j
		j := half - i
		//j越界处理
		if j > n {
			return false
		}
		if j < 0 {
			return true
		}
		//不满足左小于右处理
		if i < m && nums2[j-1] > nums1[i] {
			return false
		}
		if i > 0 && nums1[i-1] > nums2[j] {
			return true
		}
		return true
	})
	j := half - i
	var maxLeft int
	//i-1属于左，i属于右
	if i == 0 {
		maxLeft = nums2[j-1]
	} else if j == 0 {
		maxLeft = nums1[i-1]
	} else {
		maxLeft = max(nums1[i-1], nums2[j-1])
	}
	if total%2 == 0 {
		return float64(maxLeft)
	}
	var minRight int
	//i-1属于左，i属于右
	if i == m {
		minRight = nums2[j]
	} else if j == n {
		minRight = nums1[i]
	} else {
		minRight = min(nums1[i], nums1[2])
	}
	return float64(maxLeft + minRight)
}

// 三数之和
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			k := sort.Search(len(nums), func(i int) bool {
				return i > j && nums[i] >= target
			})
			if k < len(nums) && nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

// 搜索排序旋转数组
func search(nums []int, target int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		if (nums[i] >= nums[0]) == (target >= nums[0]) {
			return nums[i] >= target
		}
		if nums[i] >= nums[0] {
			return false
		}
		return true
	})
	if idx < len(nums) && nums[idx] == target {
		return idx
	}
	return -1
}

// 查找元素的第一个位置和最后一个位置
func searchRange(nums []int, target int) []int {
	start := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if start == len(nums) || (start < len(nums) && nums[start] != target) {
		return []int{-1, -1}
	}
	end := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target+1
	})
	return []int{start, end - 1}
}

// 搜索插入位置
func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}

// x的平方根
func mySqrt(x int) int {
	idx := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if idx*idx == x {
		return idx
	}
	if idx*idx > x {
		return idx - 1
	}
	return -1
}

// 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	idx := sort.Search(m*n, func(i int) bool {
		row := i / n
		col := i % n
		return matrix[row][col] >= target
	})
	if idx < m*n && matrix[idx/n][idx%n] == target {
		return true
	}
	return false
}

// 搜索旋转排序数组中的最小值
func findMin(nums []int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] <= nums[0]
	})
	if idx == len(nums) {
		return nums[0]
	}
	return nums[idx]
}

// 寻找峰值
func findPeakElement(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] > nums[i+1]
	})
}

// 搜索二维矩阵II
func searchMatrixII(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := m-1, 0
	for i >= 0 && j <= n-1 {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	var res []int
	for i := 0; i < len(nums); i++ {
		if len(res) == 0 || res[len(res)-1] < nums[i] {
			res = append(res, nums[i])
		} else {
			target := nums[i]
			idx := sort.Search(len(res), func(i int) bool {
				return res[i] >= target
			})
			res[idx] = nums[i]
		}
	}
	return len(res)
}
