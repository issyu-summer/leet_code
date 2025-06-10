package main

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"sort"
	"strconv"
	"strings"
)

func main() {
	testCases := []struct {
		input  string
		expect int
	}{
		{"eceba", 3},
		{"ccaabbb", 5},
		{"a", 1},
		{"aaaa", 4},
		{"abcabcabc", 2},
		{"", 0},
	}

	for _, tc := range testCases {
		result := lengthOfLongestSubstringTwoDistinctII(tc.input)
		fmt.Printf("Input: %-10q Expected: %d Got: %d \t %v\n",
			tc.input, tc.expect, result, result == tc.expect)
	}
}

// 零钱兑换
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 0; i < len(f); i++ {
		f[i] = amount + 1
	}
	f[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] = min(f[j], f[j-coins[i]]+1)
		}
	}
	if f[amount] == amount+1 {
		return -1
	}
	return f[amount]
}

// 最小覆盖子串
func minWindow(s string, t string) string {
	sMap := map[byte]int{}
	tMap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	check := func() bool {
		for ch, cnt := range tMap {
			if sMap[ch] < cnt {
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

// MyQueue 用栈实现队列
type MyQueue struct {
	a, b []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func (this *MyQueue) Push(x int) {
	this.a = append(this.a, x)
}

func (this *MyQueue) Pop() int {
	res := this.Peek()
	this.b = this.b[1:]
	return res
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if len(this.b) != 0 {
		return this.b[0]
	}
	for len(this.a) != 0 {
		this.b = append(this.b, this.a[0])
		this.a = this.a[1:]
	}
	return this.b[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.a) == 0 && len(this.b) == 0
}

// 组合总数
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backTrack func(path []int, start int, target int)
	backTrack = func(path []int, start int, target int) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			backTrack(append(path, candidates[i]), i, target-candidates[i])
		}
	}
	backTrack([]int{}, 0, target)
	return res
}

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	f := make([][]int, len(text1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
			}
		}
	}
	return f[len(text1)][len(text2)]
}

// 寻找峰值
func findPeakElement(nums []int) int {
	return sort.Search(len(nums)-1, func(i int) bool {
		return nums[i] >= nums[i+1]
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 环形链表II
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	exist := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			exist = true
			break
		}
	}
	if !exist {
		return nil
	}
	cur := head
	for cur != slow {
		cur = cur.Next
		slow = slow.Next
	}
	return cur
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 删除链表的第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

// 两数相加，链表版本
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	var carry int
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
		sum += carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			val := fast.Val
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

// 子集
func subsets(nums []int) [][]int {
	var res [][]int
	var backTrack func(path []int, start int)
	backTrack = func(path []int, start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			backTrack(append(path, nums[i]), i+1)
		}
	}
	backTrack([]int{}, 0)
	return res
}

// 括号生成
func generateParenthesis(n int) []string {
	var res []string
	var backTrack func(path []byte, open, close int)
	backTrack = func(path []byte, open, close int) {
		if len(path) == 2*n {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		if open < n {
			backTrack(append(path, '('), open+1, close)
		}
		if close < open {
			backTrack(append(path, ')'), open, close+1)
		}
	}
	backTrack([]byte{}, 0, 0)
	return res
}

// 求根节点到叶节点的数字之和
func sumNumbers(root *TreeNode) int {
	var helper func(root *TreeNode, prevSum int) int
	helper = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := prevSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return helper(root.Left, sum) + helper(root.Right, sum)
	}
	return helper(root, 0)
}

// 恢复IP地址
func restoreIpAddresses(s string) []string {
	var res []string
	valid := func(segment string) bool {
		if len(segment) > 1 && segment[0] == '0' {
			return false
		}
		val, err := strconv.Atoi(segment)
		if err != nil {
			return false
		}
		return val >= 0 && val <= 255
	}
	var backTrack func(path []string, remaining string, k int)
	backTrack = func(path []string, remaining string, k int) {
		if k == 4 && len(remaining) == 0 {
			res = append(res, strings.Join(append([]string{}, path...), "."))
			return
		}
		//1<=segment<=3
		for i := 1; i <= 3; i++ {
			if i > len(remaining) {
				continue
			}
			selected := remaining[:i]
			if !valid(selected) {
				continue
			}
			if k < 3 {
				if len(remaining)-i < 1*(3-k) || len(remaining)-i > 3*(3-k) {
					continue
				}
			} else {
				if len(remaining) != i {
					continue
				}
			}
			backTrack(append(path, selected), remaining[i:], k+1)
		}
	}
	backTrack([]string{}, s, 0)
	return res
}

// 排序数组
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, l, r int) int {
	randomIdx := l + rand.IntN(r-l+1)
	nums[randomIdx], nums[r] = nums[r], nums[randomIdx]
	pivot := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
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

// 字符串解码
func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			//字母出栈
			var str []byte
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				str = append([]byte{stack[len(stack)-1]}, str...)
				stack = stack[:len(stack)-1]
			}
			//]出栈
			stack = stack[:len(stack)-1]
			//数字出栈，注意这里是逆序
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

// 从前序和中序遍历构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	//pre order：root,left,right
	//in order：left,root,right
	rootVal := preorder[0]
	var rootIdx int
	for ; rootIdx < len(inorder); rootIdx++ {
		if rootVal == inorder[rootIdx] {
			break
		}
	}
	leftInOrder := inorder[:rootIdx]
	//3,9,20,15,7
	leftPreOrder := preorder[1 : 1+len(leftInOrder)]
	rightInOrder := inorder[rootIdx+1:]
	rightPreOrder := preorder[1+len(leftPreOrder):]
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(leftPreOrder, leftInOrder),
		Right: buildTree(rightPreOrder, rightInOrder),
	}
}

// 寻找2个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	total := m + n
	half := (total + 1) / 2
	if m == 0 {
		if total%2 == 0 {
			return float64(nums2[half]+nums2[half-1]) / 2
		} else {
			return float64(nums2[half-1])
		}
	}
	//左小于右
	i := sort.Search(m, func(i int) bool {
		j := half - i
		if j < 0 {
			return true
		}
		if j > n {
			return false
		}
		if i > 0 && nums1[i-1] > nums2[j] {
			return true
		}
		if i < m && nums1[i] < nums2[j-1] {
			return false
		}
		return true
	})
	j := half - i
	var leftMax int
	if i == 0 {
		leftMax = nums2[j-1]
	} else if j == 0 {
		leftMax = nums1[i-1]
	} else {
		leftMax = max(nums1[i-1], nums2[j-1])
	}
	if total%2 != 0 {
		return float64(leftMax)
	}
	var rightMin int
	if i == m {
		rightMin = nums2[j]
	} else if j == n {
		rightMin = nums1[i]
	} else {
		rightMin = min(nums1[i], nums2[j])
	}
	return float64(leftMax+rightMin) / 2
}

func multiply(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return ""
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	var ans string
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		ans += strconv.Itoa(res[i])
	}
	return ans
}

func maxSlidingWindow(nums []int, k int) []int {
	//var res []int
	//for l, r := 0, 0; r < len(nums); r++ {
	//	for l <= r && r-l+1 >= k {
	//		res = append(res, max(nums[l], nums[l+1:r+1]...))
	//		l++
	//	}
	//}
	//return res
	return []int{}
}

func maxSlidWindow(nums []int, k int) []int {
	var res, q []int //q left side
	for r := 0; r < len(nums); r++ {
		//越界
		for len(q) > 0 && r-q[0]+1 > k {
			q = q[1:]
		}
		//报纸单调递减
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		//第3个，则实际idx=2,因此r>=2时开始产生output
		if r >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

//ROOT
//Min,10,16(key)
//A,B,C(page_no)
//|					|					|
//Min,5,7				Min,10,13			Min,16,20
//D,E,F				G,H,I				J,K,L
//|		|		|
//2,3	<->	5,6	<->	7,8

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) < 3 {
		return len(s)
	}
	var res = 2
	lastIdx := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		lastIdx[s[r]]++
		for l < r && len(lastIdx) > 2 {
			lastIdx[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func lengthOfLongestSubstringTwoDistinctII(s string) int {
	if len(s) < 3 {
		return len(s)
	}
	var res = 2
	lastIdx := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		lastIdx[s[r]] = r
		if len(lastIdx) > 2 {
			var delIdx = len(s) + 1
			for _, idx := range lastIdx {
				delIdx = min(delIdx, idx)
			}
			delete(lastIdx, s[delIdx])
			l = delIdx + 1
		}
		res = min(res, r-l+1)
	}
	return res
}
