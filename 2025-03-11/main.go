package _025_03_11

import "math"

func lengthOfLastWord(s string) int {
	res := 0
	l := len(s) - 1
	for l >= 0 && s[l] == ' ' {
		l--
	}
	for i := l; i >= 0; i-- {
		res++
		if s[i] == ' ' {
			break
		}
	}
	return res
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lcp := func(a, b string) string {
		l := len(a)
		if len(a) > len(b) {
			l = len(b)
		}
		idx := 0
		for i := 0; i < l; i++ {
			if a[i] == b[i] {
				idx++
			} else {
				break
			}
		}
		return a[:idx]
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
	}
	return prefix
}

func findJudge(n int, trust [][]int) int {
	in := make([]int, n)
	out := make([]int, n)
	for _, t := range trust {
		out[t[0]-1]++
		in[t[1]-1]++
	}
	for i := 0; i < n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i + 1
		}
	}
	return -1
}

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

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	sum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for l <= r && sum >= target {
			sum -= nums[l]
			l++
			if sum == target && res > r-l+1 {
				res = r - l + 1
			}
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}
