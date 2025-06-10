package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {

}

// hashmap
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if lastIdx, ok := m[target-nums[i]]; ok {
			return []int{lastIdx, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func groupAnagrams(strs []string) [][]string {
	m := map[[128]int][]string{}
	for i := 0; i < len(strs); i++ {
		var key [128]int
		for _, b := range strs[i] {
			key[int(b-'0')]++
		}
		m[key] = append(m[key], strs[i])
	}
	var res [][]string
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

// 只搜索起点
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	res := 0
	for key := range m {
		tmp := key
		if m[tmp-1] {
			continue
		}
		tmpRes := 0
		for m[tmp] {
			tmpRes++
			tmp++
		}
		res = max(res, tmpRes)
	}
	return res
}

// 双指针
// 数字放在前面
func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func maxArea(height []int) int {
	var res = math.MinInt
	l, r := 0, len(height)-1
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			target := nums[i] + nums[j] + nums[k]
			if target == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if target < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

// 动态规划
func trap(height []int) int {
	var res int
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		if height[l] < height[r] {
			res += lMax - height[l]
			l++
		} else {
			res += rMax - height[r]
			r++
		}
	}
	return res
}

func maxSubArray(nums []int) int {
	f := make([]int, len(nums))
	var res int
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(nums[i], f[i-1]+nums[i])
	}
	return res
}

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	var res int
	duplicate := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		duplicate[s[r]]++
		for l < r && duplicate[s[r]] > 1 {
			duplicate[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func findAnagrams(s string, p string) []int {
	pMap := [26]int{}
	sMap := [26]int{}
	for i := 0; i < len(p); i++ {
		pMap[p[i]-'a']++
	}
	isAnagram := func() bool {
		return sMap == pMap
	}
	var res []int
	for l, r := 0, 0; r < len(s); r++ {
		sMap[s[r]-'a']++
		for l <= r && r-l+1 >= len(p) {
			fmt.Println(s[l : r+1])
			if isAnagram() {
				res = append(res, l)
			}
			sMap[s[l]-'a']--
			l++
		}
	}
	return res
}

func minWindow(s string, t string) string {
	sMap := map[byte]int{}
	tMap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tMap[s[i]]++
	}
	check := func() bool {
		for b, cnt := range tMap {
			if sMap[b] < cnt {
				return false
			}
		}
		return true
	}
	var res string
	for l, r := 0, 0; r < len(s); r++ {
		sMap[s[r]]++
		for l <= r && check() {
			if len(res) == 0 || len(res) > len(s[l:r+1]) {
				res = s[l : r+1]
			}
			sMap[s[l]]--
			l++
		}
	}
	return res
}

// 前缀和
func subarraySum(nums []int, k int) int {
	var res int
	var sum int
	prefix := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		//sum_j-sum_i=k,所以有了1个sum_j，如果有几个sum_i，那么就有几个和为k的子数组
		sum += nums[i]
		res += prefix[sum-k]
		prefix[sum]++
	}
	return res
}

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	//单调递减队列
	var res, q []int
	for r := 0; r < len(nums); r++ {
		//越界
		for len(q) > 0 && r-q[0]+1 > k {
			q = q[1:]
		}
		//保持单调
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		if r-0+1 >= k {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

// 模拟
func mergeIntervals(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i-1] * res[i-1]
	}
	r := 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] = r * res[j]
		r *= nums[j-1]
	}
	return res
}

func rotateNums(nums []int, k int) {
	k %= len(nums)
	reverse := func(nums []int) {
		l, r := 0, len(nums)-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// 转置+按行反转
func rotateMatrix(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	reverseHelper := func(nums []int) {
		l, r := 0, len(nums)-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	for _, nums := range matrix {
		reverseHelper(nums)
	}
}

// 原地hash
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// 矩阵
func setZeroes(matrix [][]int) {
	var usedX, usedY uint64 = 0, 0
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				usedX |= 1 << i
				usedY |= 1 << j
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if usedX&(1<<i) != 0 || usedY&(1<<j) != 0 {
				matrix[i][j] = 0
			}
		}
	}
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	l, r, t, b := 0, n-1, 0, m-1
	var res []int
	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if t <= b {
			for i := r; i >= l; i-- {
				res = append(res, matrix[b][i])
			}
			b--
		}
		if l <= r {
			for i := b; i >= t; i-- {
				res = append(res, matrix[i][l])
			}
			l++
		}
	}
	return res
}

// BST
func searchMatrixII(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i <= m-1 && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	newHead := reverseList(mid)
	a, b := head, newHead
	for b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	cur := head
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != cur {
				cur = cur.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
