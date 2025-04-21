package _025_02_24

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{Next: head}
	cur := pre
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return pre.Next
}

func removeElements1(arr []int, val int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		if arr[i] == val {
			arr[i] = arr[j]
			j--
		} else {
			i++
		}
	}
	return i
}

func removeDuplicates(arr []int) int {
	remove := func(arr []int, k int) int {
		i, j := k, k
		for j < len(arr) {
			if arr[i-k] != arr[j] {
				arr[i] = arr[j]
				i++
			}
			j++
		}
		return i
	}
	return remove(arr, 1)
}

func minimumTotal(triangle [][]int) int {
	f := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		f[i] = make([]int, len(triangle[i]))
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	return min(f[len(triangle)-1]...)
}

func min(arr ...int) int {
	res := math.MaxInt
	for _, v := range arr {
		if v < res {
			res = v
		}
	}
	return res
}

func minimumTotal1(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func majorityElement(nums []int) int {
	res := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			res = nums[i]
			cnt = 1
		}
	}
	return res
}
