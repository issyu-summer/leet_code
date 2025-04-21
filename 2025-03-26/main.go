package main

import (
	"fmt"
	"math"
)

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt64
	for l, r := 0, 0; r < len(nums); r++ {
		target -= nums[r]
		for target <= 0 && l <= r {
			if res > r-l+1 {
				res = r - l + 1
			}
			target += nums[l]
			l++
		}
	}
	if res == math.MaxInt64 {
		return 0
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]struct{}{}
	res := 0
	for l, r := 0, 0; r < len(s); r++ {
		// l和r是否相等，与check有关，，在该例子中，l==r时，check一定成立，会导致l越过r
		_, ok := m[s[r]]
		for ok && l < r {
			delete(m, s[l])
			l++
			_, ok = m[s[r]]
		}
		res = min(res, r-l+1)
		m[s[r]] = struct{}{}
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}
	res := make([][]string, 0)
	for _, strArr := range m {
		fmt.Println(strArr)
		res = append(res, strArr)
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	a, b := 0, 0
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += a
		a = sum / 10
		b = sum % 10
		cur.Next = &ListNode{
			Val: b,
		}
		cur = cur.Next
	}
	if a != 0 {
		cur.Next = &ListNode{
			Val: a,
		}
	}
	return pre.Next
}
