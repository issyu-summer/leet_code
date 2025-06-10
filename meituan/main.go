package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sort"
	"strconv"
	"strings"
)

func main() {

}

// 合并2个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	fmt.Println(i, j, k)
	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// ListNode 反转链表2
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	lPre, l, r, rNext := dummy, dummy, dummy, dummy
	for i := 0; i < left-1; i++ {
		lPre = lPre.Next
	}
	l = lPre.Next
	for i := 0; i < right; i++ {
		r = r.Next
	}
	rNext = r.Next
	_ = reverseHelper(l, rNext)
	lPre.Next = r
	l.Next = rNext
	return dummy.Next
}

func reverseHelper(start, end *ListNode) *ListNode {
	var pre, cur *ListNode = nil, start
	for cur != end {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// 环形链表
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

type Node struct {
	key   int
	val   int
	left  *Node
	right *Node
}
type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	this := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	this.head.right = this.tail
	this.tail.left = this.head
	return this
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.addToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.remove(node)
		this.addToHead(node)
		return
	}
	if len(this.cache) >= this.capacity {
		back := this.tail.left
		this.remove(back)
		delete(this.cache, back.key)
	}
	node := &Node{key: key, val: value}
	this.addToHead(node)
	this.cache[key] = node
}

func (this *LRUCache) remove(node *Node) {
	first, _, third := node.left, node, node.right
	first.right = third
	third.left = first
}

func (this *LRUCache) addToHead(node *Node) {
	first, second, third := this.head, node, this.head.right

	first.right = second
	second.left = first

	second.right = third
	third.left = second
}

// 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	l, r := dummy, dummy.Next
	for r != nil && r.Next != nil {
		if r.Val == r.Next.Val {
			val := r.Val
			for r != nil && r.Val == val {
				r = r.Next
			}
			l.Next = r
		} else {
			r = r.Next
			l = l.Next
		}
	}
	return dummy.Next
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
		//segment最长为3,最短为1
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

// 重排链表
func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	midNxt := mid.Next
	mid.Next = nil
	newHead := reverseHelper(midNxt, nil)
	first, second := head, newHead
	for second != nil {
		firstNxt, secondNxt := first.Next, second.Next
		first.Next = second
		second.Next = firstNxt

		first = firstNxt
		second = secondNxt
	}
}

// 有序数组中的重复项
func removeDuplicates(nums []int) int {
	removeHelp := func(nums []int, k int) int {
		i, j := k, k
		for j < len(nums) {
			if nums[j] != nums[i-k] {
				nums[i] = nums[j]
				i++
			}
			j++
		}
		return i
	}
	return removeHelp(nums, 1)
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var (
		res    = 0
		bondry = math.MaxInt32 / 10
		last   = math.MaxInt32 % 10
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
		if res > bondry || res == bondry && int(s[i]-'0') > last {
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

func sortArray(nums []int) []int {
	var partition func(nums []int, l, r int) int
	partition = func(nums []int, l, r int) int {
		randIdx := l + rand.IntN(r-l+1)
		nums[r], nums[randIdx] = nums[randIdx], nums[r]
		i := l
		pivot := nums[r]
		for j := l; j < r; j++ {
			if nums[j] < pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[r] = nums[r], nums[i]
		return i
	}
	var sort func(nums []int, l, r int)
	sort = func(nums []int, l, r int) {
		if l >= r {
			return
		}
		pivotIdx := partition(nums, l, r)
		sort(nums, l, pivotIdx-1)
		sort(nums, pivotIdx+1, r)
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var helper func(root *TreeNode, preSum int) int
	helper = func(root *TreeNode, preSum int) int {
		if root == nil {
			return 0
		}
		sum := preSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return helper(root.Left, sum) + helper(root.Right, sum)
	}
	return helper(root, 0)
}

func addStrings(num1 string, num2 string) string {
	var (
		res   []byte
		carry int
	)
	var i, j = len(num1) - 1, len(num2) - 1
	for i >= 0 || j >= 0 {
		var sum int
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		sum += carry
		carry = sum / 10
		sum = sum % 10
		fmt.Println(sum, carry)
		res = append([]byte{byte(sum + '0')}, res...)
	}
	fmt.Println(carry)
	if carry > 0 {
		res = append([]byte{byte(carry + '0')}, res...)
	}
	return string(res)
}

func isValid(s string) bool {
	stack := []byte{}
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' || len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if m[stack[len(stack)-1]] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func maxPathSum(root *TreeNode) int {
	var res int = math.MinInt
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(maxGain(root.Left), 0)
		r := max(maxGain(root.Right), 0)
		res = max(res, l+r+root.Val)
		return max(l, r) + root.Val
	}
	maxGain(root)
	return res
}

func peakIndexInMountainArray(arr []int) int {
	i, j := 0, len(arr)
	for i < j {
		h := (i + j) >> 1
		if i < len(arr)-1 && arr[h] < arr[h+1] {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func deleteDuplicatesInList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
	for fast != nil && fast.Next != nil {
		if fast.Val == fast.Next.Val {
			val := fast.Val
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

// abab
// bababa
func repeatedSubstringPattern(s string) bool {
	doubled := s + s
	return strings.Contains(doubled[1:len(doubled)-1], s)
}

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minD, minDepth(root.Left))
	}
	if root.Right != nil {
		minD = min(minD, minDepth(root.Right))
	}
	return minD + 1
}

// 二叉树的完全性校验
func isCompleteTree(root *TreeNode) bool {
	var end bool
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			end = true
			continue
		}
		if end {
			return false
		}
		q = append(q, node.Left)
		q = append(q, node.Right)
	}
	return true
}

func removeDuplicatesInNums(nums []int) int {
	remove := func(nums []int, k int) int {
		i, j := k, k
		for j < len(nums) {
			if nums[j] != nums[j-k] {
				nums[i] = nums[j]
				i++
			}
			j++
		}
		return i
	}
	return remove(nums, 1)
}

// 二叉树的路径和
func pathTarget(root *TreeNode, target int) [][]int {
	var res [][]int
	var backTrack func(root *TreeNode, path []int, target int)
	backTrack = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && target == root.Val {
			res = append(res, append(append([]int{}, path...), root.Val))
			return
		}
		backTrack(root.Left, append(path, root.Val), target-root.Val)
		backTrack(root.Right, append(path, root.Val), target-root.Val)
	}
	backTrack(root, []int{}, target)
	return res
}

// 打家劫舍III
func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		left, noLeft := dfs(root.Left)
		right, noRight := dfs(root.Right)
		cur := root.Val + noLeft + noRight
		noCur := max(left, noLeft) + max(right, noRight)
		return cur, noCur
	}
	return max(dfs(root))
}

// 打家劫舍II
func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	robLinar := func(nums []int) int {
		f := make([]int, len(nums))
		f[0] = nums[0]
		f[1] = max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			f[i] = max(f[i-2]+nums[i], f[i-1])
		}
		return f[len(nums)-1]
	}
	return max(robLinar(nums[:len(nums)-1]), robLinar(nums[1:]))
}

func inventoryManagement(stock []int, k int) []int {
	if k <= 0 || len(stock) == 0 {
		return nil
	}

	// 前k个元素构建最大堆
	heap := make([]int, k)
	copy(heap, stock[:k])
	buildMaxHeap(heap)

	// 处理剩余元素
	for i := k; i < len(stock); i++ {
		if stock[i] < heap[0] {
			heap[0] = stock[i]
			maxHeapify(heap, 0, k)
		}
	}

	return heap
}

// 构建最大堆
func buildMaxHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, n)
	}
}

// 维护最大堆性质
func maxHeapify(arr []int, i, heapSize int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}

	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest, heapSize)
	}
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	var dirs = []struct{ x, y int }{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if nextRow, nextCol := row+dir.x, col+dir.y; nextRow < 0 || nextRow >= n ||
			nextCol < 0 || nextCol >= n || matrix[nextRow][nextCol] > 0 {
			dirIdx = (dirIdx + 1) % 4
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}

func singleNonDuplicate(nums []int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		if i%2 == 1 {
			i-- // 确保总是检查偶数索引
		}
		return i == len(nums)-1 || nums[i] != nums[i+1]
	})
	if idx%2 == 1 {
		idx--
	}
	return nums[idx]
}
