package main

import (
	"container/list"
	"fmt"
	"math"
	"strings"
)

func main() {

}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entry).val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	//1.if exist update
	if elem, ok := l.cache[key]; ok {
		elem.Value.(*entry).val = value
		l.list.MoveToFront(elem)
		return
	}
	//2.if len >= capacity evict back
	if l.list.Len() >= l.capacity {
		back := l.list.Back()
		l.list.Remove(back)
		delete(l.cache, back.Value.(*entry).key)
	}
	//3.put in cache
	elem := l.list.PushFront(&entry{key, value})
	l.cache[key] = elem
}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		//save left side
		q   []int
		res []int
	)
	//r-l+1=k => l=r-k+1
	for r := 0; r < len(nums); r++ {
		//左边姐越界
		for len(q) > 0 && q[0] < r-k+1 {
			q = q[1:]
		}
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		if r >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var (
		res   = 0
		bndry = math.MaxInt32 / 10
		last  = math.MaxInt32 % 10
	)
	var (
		idx  = 0
		sign = 1
	)
	if s[0] == '+' {
		idx = 1
		sign = 1
	}
	if s[0] == '-' {
		idx = 1
		sign = -1
	}
	for i := idx; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		if res > bndry || (res == bndry && int(s[i]-'0') > last) {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + int(s[i]-'0')
	}
	return res * sign
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	var (
		sum, carry int
	)
	for l1 != nil || l2 != nil {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		sum = sum % 10
		fmt.Println(sum, carry)
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func longestValidParentheses(s string) int {
	var (
		stack []int
		res   int
	)
	for r := 0; r < len(s); r++ {
		if s[r] == '(' {
			stack = append(stack, r)
		} else {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				var curLen int
				if len(stack) == 0 {
					curLen = r + 1
				} else {
					//必须取上一个idx，才是全部的长度，否则要累加
					curLen = r - stack[len(stack)-1]
				}
				res = max(res, curLen)
			} else {
				stack = append(stack, r)
			}
		}
	}
	return res
}
