package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"math/rand/v2"
	"sort"
)

func main() {
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	nums := []int{1, 3, 8, 5, 1, 9, 8}
	quickSort(nums, 0, 6)
	fmt.Println(nums)
}

// codetop *4

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.K个一组反转链表
func reverseKGroupAll(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	cur := head
	for i := 0; i < k && cur != nil; i++ {
		cur = cur.Next
	}
	newHead := reverse(head, cur)
	head.Next = reverseKGroupAll(cur, k)
	return newHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		//小于k个的链表，长度不变
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	//先处理K个反转
	newHead := reverse(head, cur)
	//然后分治K之后的翻转
	head.Next = reverseKGroup(cur, k)
	return newHead
}

func reverse(head *ListNode, end *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != end {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// 2.三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	//n-2至少3位数，n-1,n-2,n-3
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//使用二分搜索，查找l,r，满足条件三数之和为0
		for j := i + 1; j < n-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			//选择i
			//loop选择l
			target := 0 - nums[i] - nums[j]
			fmt.Println("search", target)
			//实际从j+1后的第一个位置开始搜索，则k在nums中的idx为+j+1
			fmt.Println("search window", nums[j+1:])
			k := sort.Search(n, func(idx int) bool {
				return idx > j && nums[idx] >= target
			})
			fmt.Println("i,j,k", i, j, k)
			if k < n {
				fmt.Println(nums[i], nums[j], nums[k])
			}
			if k < n && nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

func threeSumTwoP(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	//n-2至少3位数，n-1,n-2,n-3
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				//skip duplicate
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// 3.最大子数组和
func maxSubArray(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		if res < f[i] {
			res = f[i]
		}
	}
	return res
}

type Ints []int

func (h *Ints) Len() int {
	return len(*h)
}

func (h *Ints) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Ints) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Ints) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Ints) Pop() any {
	old := *h
	n := len(old)
	x := (*h)[n-1]
	*h = old[:n-1]
	return x
}

// 4.快速排序
func partition(nums []int, l, r int) int {
	//随机基准
	if r == 6 {
		fmt.Println(nums)
	}
	randIdx := l + rand.IntN(r-l+1)
	nums[r], nums[randIdx] = nums[randIdx], nums[r]
	//分区
	pivot := nums[r]
	if r == 6 {
		fmt.Println(nums, pivot)
	}
	i := l
	for j := l; j < r; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		if r == 6 {
			fmt.Println(nums, pivot)
		}
	}
	//i是第pivot的插入位置，比i小的idx的元素，都小于pivot
	nums[i], nums[r] = nums[r], nums[i]
	fmt.Println(nums, pivot)
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

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// leetcode hot 100

// 1.数组中第k大元素
func findKthLargest(nums []int, k int) int {
	h := new(Ints)
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func findKthLargestWithPartition(nums []int, k int) int {
	partition := func(nums []int, l, r int) int {
		pivot := nums[r]
		i := l
		for j := l; j < r; j++ {
			if nums[j] <= pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[r] = nums[r], nums[i]
		return i
	}
	target := len(nums) - k
	l, r := 0, len(nums)-1
	for {
		idx := partition(nums, l, r)
		if idx == target {
			return nums[idx]
		} else if idx < target {
			l = idx + 1
		} else {
			r = idx - 1
		}
	}
}

// 2.字符串解码
func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			//暂存,用于后续重复字符串
			var str []byte
			//非[字符
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				//先进后出
				str = append([]byte{stack[len(stack)-1]}, str...)
				stack = stack[:len(stack)-1]
			}
			//[字符，出栈，并重复stack
			stack = stack[:len(stack)-1]
			num := 0
			base := 1
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				num += int(stack[len(stack)-1]-'0') * base
				stack = stack[:len(stack)-1]
				base *= 10
			}
			stack = append(stack, bytes.Repeat(str, num)...)
		}
	}
	return string(stack)
}

// 3.找打重复数字
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
		fmt.Println("slow", slow)
		fmt.Println("fast", fast)
		fmt.Println("nums of slow idx", nums[slow])
		fmt.Println("nums of fast idx", nums[nums[fast]])
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
