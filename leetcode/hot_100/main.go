package main

import (
	"bytes"
	"container/heap"
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	var carry int
	for l1 != nil || l2 != nil {
		var sum = carry
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
		cur.Next = &ListNode{sum, nil}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
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

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, dummy.Next
	for cur != nil && cur.Next != nil {
		first := cur
		second := first.Next
		third := second.Next

		pre.Next = second
		second.Next = first
		first.Next = third

		pre = first
		cur = third
	}
	return dummy.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	end := head
	for i := 0; i < k; i++ {
		if end == nil {
			return head
		}
		end = end.Next
	}
	//end-1->head
	newHead := reverseHelper(head, end)
	head.Next = reverseKGroup(end, k)
	return newHead
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

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	newList := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		newNode := cur.Next
		cur.Next = newNode.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}
	return newList
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(mid)
	return mergeTwoLists(l1, l2)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var merge func(lists []*ListNode, l, r int) *ListNode
	merge = func(lists []*ListNode, l, r int) *ListNode {
		if l >= r {
			return lists[l]
		}
		mid := (l + r) >> 1
		l1 := merge(lists, l, mid)
		l2 := merge(lists, mid+1, r)
		return mergeTwoLists(l1, l2)
	}
	return merge(lists, 0, len(lists)-1)
}

type Entry struct {
	key, val    int
	left, right *Entry
}
type LRUCache struct {
	capacity   int
	cache      map[int]*Entry
	head, tail *Entry
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    map[int]*Entry{},
		head:     &Entry{},
		tail:     &Entry{},
	}
	lru.head.right = lru.tail
	lru.tail.left = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if entry, ok := lru.cache[key]; ok {
		lru.remove(entry)
		lru.addToHead(entry)
		return entry.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, val int) {
	if entry, ok := lru.cache[key]; ok {
		entry.val = val
		lru.remove(entry)
		lru.addToHead(entry)
		return
	}
	if len(lru.cache) >= lru.capacity {
		end := lru.tail.left
		lru.remove(end)
		delete(lru.cache, end.key)
	}
	entry := &Entry{key, val, nil, nil}
	lru.cache[key] = entry
	lru.addToHead(entry)
}

func (lru *LRUCache) remove(entry *Entry) {
	first, _, third := entry.left, entry, entry.right
	first.right = third
	third.left = first
}

func (lru *LRUCache) addToHead(entry *Entry) {
	first, second, third := lru.head, entry, lru.head.right

	first.right = second
	second.left = first

	second.right = third
	third.left = second
}

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var inorder func(*TreeNode, func(node *TreeNode))
	var res []int
	inorder = func(root *TreeNode, f func(node *TreeNode)) {
		if root == nil {
			return
		}
		inorder(root.Left, f)
		f(root)
		inorder(root.Right, f)
	}
	inorder(root, func(node *TreeNode) {
		res = append(res, node.Val)
	})
	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return max(l, r) + 1
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func invertTreeBFS(root *TreeNode) *TreeNode {
	var bfs func(root *TreeNode, handler func([]*TreeNode))
	bfs = func(root *TreeNode, handler func([]*TreeNode)) {
		if root == nil {
			return
		}
		q := []*TreeNode{root}
		for len(q) > 0 {
			var level []*TreeNode
			levelSize := len(q)
			for i := 0; i < levelSize; i++ {
				node := q[0]
				q = q[1:]
				level = append(level, node)
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
			handler(level)
		}
	}
	bfs(root, func(level []*TreeNode) {
		for i := 0; i < len(level); i++ {
			level[i].Left, level[i].Right = level[i].Right, level[i].Left
		}
	})
	return root
}

func isSymmetric(root *TreeNode) bool {
	var check func(l, r *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil || l.Val != r.Val {
			return false
		}
		return check(l.Left, r.Right) && check(l.Right, r.Left)
	}
	return check(root.Left, root.Right)
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := depth(root.Left)
		r := depth(root.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}
	depth(root)
	return res
}

func levelOrder(root *TreeNode) [][]int {
	var bfs func(root *TreeNode, handler func([]*TreeNode))
	bfs = func(root *TreeNode, handler func([]*TreeNode)) {
		if root == nil {
			return
		}
		q := []*TreeNode{root}
		for len(q) > 0 {
			var level []*TreeNode
			levelSize := len(q)
			for i := 0; i < levelSize; i++ {
				node := q[0]
				q = q[1:]
				level = append(level, node)
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
			handler(level)
		}
	}
	var res [][]int
	bfs(root, func(level []*TreeNode) {
		var ints []int
		for i := 0; i < len(level); i++ {
			ints = append(ints, level[i].Val)
		}
		res = append(res, ints)
	})
	return res
}

// 左->根->右
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(nums []int, l, r int) *TreeNode
	helper = func(nums []int, l, r int) *TreeNode {
		//l>=r，仍然还有一个节点，所以不能跳过
		if l > r {
			return nil
		}
		mid := (l + r) >> 1
		return &TreeNode{
			Val:   nums[mid],
			Left:  helper(nums, l, mid-1),
			Right: helper(nums, mid+1, r),
		}
	}
	return helper(nums, 0, len(nums)-1)
}

func isValidBST(root *TreeNode) bool {
	var check func(root *TreeNode, lower, upper int) bool
	check = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		return check(root.Left, lower, root.Val) && check(root.Right, root.Val, upper)
	}
	return check(root, math.MinInt, math.MaxInt)
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var inorder func(root *TreeNode, handler func(node *TreeNode))
	inorder = func(root *TreeNode, handler func(node *TreeNode)) {
		if root == nil {
			return
		}
		inorder(root.Left, handler)
		k--
		if k == 0 {
			handler(root)
		}
		inorder(root.Right, handler)
	}
	inorder(root, func(node *TreeNode) {
		res = node.Val
	})
	return res
}

func rightSideView(root *TreeNode) []int {
	var bfs func(root *TreeNode, handler func(level []*TreeNode))
	bfs = func(root *TreeNode, handler func(level []*TreeNode)) {
		if root == nil {
			return
		}
		q := []*TreeNode{root}
		for len(q) > 0 {
			var level []*TreeNode
			var levelSize = len(q)
			for i := 0; i < levelSize; i++ {
				node := q[0]
				q = q[1:]
				level = append(level, node)
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
			handler(level)
		}
	}
	var res []int
	bfs(root, func(level []*TreeNode) {
		res = append(res, level[len(level)-1].Val)
	})
	return res
}

func rightSideViewII(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode, depth int, handler func(node *TreeNode))
	dfs = func(root *TreeNode, depth int, handler func(node *TreeNode)) {
		if root == nil {
			return
		}
		if len(res) == depth {
			handler(root)
		}
		dfs(root.Right, depth+1, handler)
		dfs(root.Left, depth+1, handler)
	}
	dfs(root, 0, func(node *TreeNode) {
		res = append(res, node.Val)
	})
	return res
}

func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			nxt := cur.Left
			pre := nxt
			if nxt.Right != nil {
				pre = nxt.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, nxt
		}
		cur = cur.Right
	}
}

func flattenII(root *TreeNode) {
	var inorder func(*TreeNode, func(node *TreeNode))
	inorder = func(root *TreeNode, f func(node *TreeNode)) {
		if root == nil {
			return
		}
		f(root)
		inorder(root.Left, f)
		inorder(root.Right, f)
	}
	var lists []*TreeNode
	inorder(root, func(node *TreeNode) {
		lists = append(lists, node)
	})
	for i := 1; i < len(lists); i++ {
		pre, cur := lists[i-1], lists[i]
		pre.Left, pre.Right = nil, cur
	}
}

//pre:root->l->r
//in:l->root->
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	var inorderIdx int
	for inorderIdx < len(inorder) {
		if inorder[inorderIdx] == val {
			break
		}
		inorderIdx++
	}
	lIn, _, rIn := inorder[:inorderIdx], inorder[inorderIdx], inorder[inorderIdx+1:]
	_, lPre, rPre := preorder[0], preorder[1:len(lIn)+1], preorder[len(lIn)+1:]
	return &TreeNode{
		Val:   val,
		Left:  buildTree(lPre, lIn),
		Right: buildTree(rPre, rIn),
	}
}

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	prefix := map[int]int{0: 1}
	var backTrack func(root *TreeNode, sum int)
	backTrack = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		res += prefix[sum+root.Val-targetSum]
		prefix[sum+root.Val]++
		backTrack(root.Left, sum+root.Val)
		backTrack(root.Right, sum+root.Val)
		prefix[sum+root.Val]--
	}
	backTrack(root, 0)
	return res
}

func pathSumIII(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var backTrack func(root *TreeNode, target int) int
	backTrack = func(root *TreeNode, target int) int {
		if root == nil {
			return 0
		}
		cnt := 0
		if target == root.Val {
			cnt = 1
		}
		return cnt + backTrack(root.Left, target-root.Val) + backTrack(root.Right, target-root.Val)
	}
	var preorder func(root *TreeNode) int
	preorder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return backTrack(root, targetSum) + preorder(root.Left) + preorder(root.Right)
	}
	return preorder(root)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(0, maxGain(root.Left))
		r := max(0, maxGain(root.Right))
		res = max(res, r+l+root.Val)
		return max(l, r) + root.Val
	}
	maxGain(root)
	return res
}

// 图
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	var (
		fresh, minutes int
		q              [][]int
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	directions := [][]int{{-1, 0}, {1, 0}, {-1, 0}, {1, 0}}
	for len(q) > 0 {
		levelSize := len(q)
		rottenInThisLeve := false
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			for _, dir := range directions {
				x, y := cur[0]+dir[0], cur[1]+dir[1]
				if x < 0 || x > m-1 || y < 0 || y > n-1 || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 2
				fresh--
				q = append(q, []int{x, y})
				rottenInThisLeve = true
			}
		}
		if rottenInThisLeve {
			minutes++
		}
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int, numCourses)
	inDegree := make([]int, numCourses)
	//1->[0,2]
	//2->[1]
	//0->[2]
	//3->[0]
	for i := 0; i < len(prerequisites); i++ {
		//a->b
		a, b := prerequisites[i][1], prerequisites[i][0]
		adj[a] = append(adj[a], b)
		inDegree[b]++
	}
	var q []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	cnt := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		cnt++
		bList := adj[cur]
		for i := 0; i < len(bList); i++ {
			b := bList[i]
			inDegree[b]--
			if inDegree[b] == 0 {
				q = append(q, b)
			}
		}
	}
	return cnt == numCourses
}

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func ConstructorTrie() Trie {
	return Trie{
		children: map[rune]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &Trie{children: map[rune]*Trie{}}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

// 回溯
func permute(nums []int) [][]int {
	var res [][]int
	var fc func(path []int, used uint64)
	fc = func(path []int, used uint64) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			masking := uint64(1 << i)
			if used&masking != 0 {
				continue
			}
			fc(append(path, nums[i]), used|masking)
		}
	}
	fc([]int{}, 0)
	return res
}

// 电话号码的组合
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var res []string
	var fc func(path []byte, start int)
	fc = func(path []byte, start int) {
		if len(digits) == len(path) {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		bytes := m[digits[start]]
		for i := 0; i < len(bytes); i++ {
			fc(append(path, bytes[i]), start+1)
		}
	}
	fc([]byte{}, 0)
	return res
}

// 子集
func subsets(nums []int) [][]int {
	var res [][]int
	var fc func(path []int, start int)
	fc = func(path []int, start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			fc(append(path, nums[i]), i+1)
		}
	}
	fc([]int{}, 0)
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var fc func(path []int, start, target int)
	fc = func(path []int, start, target int) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if target < 0 {
			return
		}
		//不是每次都从头开始，否则res会有重复
		for i := start; i < len(candidates); i++ {
			fc(append(path, candidates[i]), i, target-candidates[i])
		}
	}
	fc([]int{}, 0, target)
	return res
}

func generateParenthesis(n int) []string {
	var res []string
	var fc func(path []byte, open, close int)
	fc = func(path []byte, open, close int) {
		if len(path) == n*2 {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		if open < n {
			fc(append(path, '('), open+1, close)
		}
		if close < open {
			fc(append(path, ')'), open, close+1)
		}
	}
	fc([]byte{}, 0, 0)
	return res
}

func exist(board [][]byte, word string) bool {
	var dfs func(board [][]byte, i, j int, idx int) bool
	dfs = func(board [][]byte, i, j int, idx int) bool {
		if idx == len(word) {
			return true
		}
		if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[idx] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '#'
		if dfs(board, i+1, j, idx+1) ||
			dfs(board, i-1, j, idx+1) ||
			dfs(board, i, j+1, idx+1) ||
			dfs(board, i, j-1, idx+1) {
			return true
		}
		board[i][j] = tmp
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func partition(s string) [][]string {
	ok := func(str string) bool {
		l, r := 0, len(str)-1
		for l < r {
			if str[l] != str[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	var res [][]string
	var fc func(path []string, start int)
	fc = func(path []string, start int) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			if ok(s[start : i+1]) {
				//取到i，然后从i+1继续回溯
				fc(append(path, s[start:i+1]), i+1)
			}
		}
	}
	fc([]string{}, 0)
	return res
}

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}

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

func findMin(nums []int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] < nums[0]
	})
	if idx < len(nums) {
		return nums[idx]
	}
	return nums[0]
}

// 栈
func isValid(s string) bool {
	var stack []int
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 && m[s[stack[len(stack)-1]]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// MinStack push,pop都一起，push最小值等于栈顶和val二者中较小的一个
type MinStack struct {
	stack []int
	minus []int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		stack: []int{},
		minus: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.minus[len(this.minus)-1]
	this.minus = append(this.minus, min(top, val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minus = this.minus[:len(this.minus)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minus[len(this.minus)-1]
}

func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			var str []byte
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				str = append([]byte{stack[len(stack)-1]}, str...)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			num, base := 0, 1
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

// 每日温度，单调递减栈
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIdx := stack[len(stack)-1]
			res[prevIdx] = i - prevIdx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

// 柱状图中的最大矩形，单调递增栈+尝试向右扩展
func largestRectangleArea(heights []int) int {
	var res int
	stack := []int{}
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//计算top位置及以前的最大值
			height := heights[top]
			var width int
			if len(stack) == 0 {
				//右边界-(左边界+1)
				width = i - (-1 + 1)
			} else {
				//1,4,1
				//右边界-(左边界+1)
				width = i - (stack[len(stack)-1] + 1)
			}
			res = max(res, height*width)
		}
		stack = append(stack, i)
	}
	return res
}

func findKthLargest(nums []int, k int) int {
	p := func(nums []int, l, r int) int {
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
	l, r := 0, len(nums)-1
	target := len(nums) - k
	for {
		idx := p(nums, l, r)
		if idx == target {
			return nums[idx]
		} else if idx > target {
			r = idx - 1
		} else {
			l = idx + 1
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var unique []int
	for key := range m {
		unique = append(unique, key)
	}
	p := func(nums []int, l, r int, priority func(int) int) int {
		pivot := priority(nums[r])
		i := l
		for j := l; j < r; j++ {
			if priority(nums[j]) > pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[r] = nums[r], nums[i]
		return i
	}
	var quickSelect func(nums []int, l, r, k int)
	quickSelect = func(nums []int, l, r, k int) {
		if l >= r {
			return
		}
		pivotIdx := p(nums, l, r, func(num int) int {
			return m[num]
		})
		if pivotIdx == k-1 {
			return
		} else if pivotIdx < k-1 {
			quickSelect(nums, pivotIdx+1, r, k)
		} else {
			quickSelect(nums, l, pivotIdx-1, k)
		}
	}
	quickSelect(unique, 0, len(unique)-1, k)
	return unique[:k]
}

type (
	MaxHeap []int
	MinHeap []int
)

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

func ConstructorMedianFinder() MedianFinder {
	return MedianFinder{
		maxHeap: &MaxHeap{},
		minHeap: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 决定插入哪个堆,小的放在maxHeap,大的放在minHeap
	if this.maxHeap.Len() == 0 || num <= (*this.maxHeap)[0] {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}

	// 平衡堆的大小,保证maxHeap>=minHeap+1,max一定比min多
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	} else if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	//奇数个
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	}
	//偶数个
	return float64((*this.maxHeap)[0]+(*this.minHeap)[0]) / 2.0
}

func canJump(nums []int) bool {
	var maxPosition int
	for i := 0; i < len(nums); i++ {
		if i > maxPosition {
			return false
		}
		curEnd := i + nums[i]
		maxPosition = max(maxPosition, curEnd)
		if maxPosition >= len(nums)-1 {
			return true
		}
	}
	return false
}

func jump(nums []int) int {
	cnt, end, maxPosition := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		maxPosition = max(maxPosition, i+nums[i])
		//必须要跳了
		if end == i {
			cnt++
			end = maxPosition
		}
		if end >= len(nums)-1 {
			return cnt
		}
	}
	return 0
}

func partitionLabels(s string) []int {
	lastPos := map[byte]int{}
	for i := 0; i < len(s); i++ {
		lastPos[s[i]] = i
	}
	var res []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if lastPos[s[i]] > end {
			end = lastPos[s[i]]
		}
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}

func rob(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	f[1] = max(f[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f[i] = max(f[i-2]+nums[i], f[i-1])
	}
	return f[len(nums)-1]
}

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = n + 1
	}
	f[0] = 0
	f[1] = f[1-1*1] + 1
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 0; i < len(f); i++ {
		f[i] = amount + 1
	}
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

func wordBreak(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if f[j] && dict[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}
func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	f := make([]bool, target+1)
	f[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			f[j] = f[j] || f[j-nums[i]]
		}
	}
	return f[target]
}

func longestValidParentheses(s string) int {
	var stack []int
	var res int
	for r := 0; r < len(s); r++ {
		if s[r] == '(' {
			stack = append(stack, r)
		} else {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				var curLen int
				if len(stack) == 0 {
					curLen = r - 0 + 1
				} else {
					curLen = r - (stack[len(stack)-1] + 1) + 1
				}
				res = max(res, curLen)
			} else {
				stack = append(stack, r)
			}
		}
	}
	return res
}
