package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func main() {

}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	f := make([]int, len(nums))
	f[0] = nums[0]
	f[1] = max(f[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f[i] = max(f[i-2]+nums[i], f[i-1])
	}
	return f[len(nums)-1]
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		return a+b > b+a
	})
	var res string
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	if res[0] == '0' {
		res = "0"
	}
	return res
}

func subarraySum(nums []int, k int) int {
	prefix := map[int]int{0: 1}
	var cnt int
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		cnt += prefix[sum-k]
		prefix[sum]++
	}
	return cnt
}

func subarraySumII(nums []int, k int) [][]int {
	prefix := map[int][]int{0: {-1}}
	var res [][]int
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if indicates, ok := prefix[sum-k]; ok {
			for _, start := range indicates {
				res = append(res, nums[start+1:i+1])
			}
		}
		prefix[sum] = append(prefix[sum], i)
	}
	return res
}

func majorityElement(nums []int) int {
	var res, cnt int
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
		}
		if res == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

func calculate(s string) int {
	var stack []int
	var num int
	var sign byte = '+'
	//33+2*2
	for i, ch := range s {
		isDigit := unicode.IsDigit(ch)
		fmt.Println(ch, isDigit)
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if (!isDigit && ch != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = s[i]
			num = 0
		}
	}
	fmt.Println(num, sign)
	var sum int
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}
	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
	for fast != nil {
		if fast.Next != nil && fast.Next.Val == fast.Val {
			val := fast.Val
			//第一个不等于val的位置(最后一个等于val的位置的下一个位置)
			for fast != nil && fast.Val == val {
				fast = fast.Next
			}
			slow.Next = fast
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return dummy.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
	for fast != nil {
		if fast.Next != nil && fast.Next.Val == fast.Val {
			val := fast.Val
			//最后一个等于val的位置
			for fast.Next != nil && fast.Next.Val == val {
				fast = fast.Next
			}
			slow.Next = fast
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return dummy.Next
}
