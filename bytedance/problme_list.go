package main

import (
	"math"
	"math/rand/v2"
	"slices"
	"sort"
)

func main() {
	m := make(map[int64]int32)
	m[1] = 2
}

// 无重复最长子串
func lengthOfLongestSubstring(s string) int {
	var res int
	m := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		m[s[r]]++
		for l < r && m[s[r]] > 1 {
			m[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

// k个一组反转链表
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
	key, val    int
	left, right *Node
}
type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func ConstructorLRUCache(capacity int) LRUCache {
	lruCache := LRUCache{
		capacity: capacity,
		cache:    map[int]*Node{},
		head:     &Node{},
		tail:     &Node{},
	}
	lruCache.head.right = lruCache.tail
	lruCache.tail.left = lruCache.head
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.remove(elem)
		this.addToHead(elem)
		return elem.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		elem.val = value
		this.remove(elem)
		this.addToHead(elem)
		return
	}
	if len(this.cache) >= this.capacity {
		end := this.tail.left
		this.remove(end)
		delete(this.cache, end.key)
	}
	node := &Node{key, value, nil, nil}
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

func findKthLargest(nums []int, k int) int {
	p := func(nums []int, l, r int) int {
		randIdx := l + rand.IntN(r-l+1)
		nums[randIdx], nums[r] = nums[r], nums[randIdx]
		var (
			i     = l
			pivot = nums[r]
		)
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
		} else if idx < target {
			l = idx + 1
		} else {
			r = idx - 1
		}
	}
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	q := []*TreeNode{root}
	var LTR = true
	for len(q) > 0 {
		levelSize := len(q)
		var level []int
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if LTR {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		LTR = !LTR
		res = append(res, level)
	}
	return res
}

func numIslands(grid [][]byte) int {
	var res int
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[i])-1 || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func search(nums []int, target int) int {
	check := func(mid int) bool {
		if (nums[mid] >= nums[0]) == (target >= nums[0]) {
			return nums[mid] >= target
		}
		if nums[mid] >= nums[0] {
			return false
		}
		return true
	}
	i, j := 0, len(nums)
	for i < j {
		mid := (i + j) >> 1
		if check(mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if i < len(nums) && nums[i] == target {
		return i
	}
	return -1
}

func maxProfit(prices []int) int {
	var res int
	var cost = prices[0]
	for i := 1; i < len(prices); i++ {
		cost = min(cost, prices[i])
		res = max(res, prices[i]-cost)
	}
	return res
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	m, n := len(matrix), len(matrix[0])
	l, r, t, b := 0, n-1, 0, m-1
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

func trap(height []int) int {
	n := len(height)
	lMax, rMax := make([]int, n), make([]int, n)
	lMax[0], rMax[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		lMax[i] = max(height[i], lMax[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(height[i], rMax[i+1])
	}
	var res int
	for i := 0; i < n; i++ {
		res += min(lMax[i], rMax[i]) - height[i]
	}
	return res
}

func longestPalindrome(s string) string {
	expand := func(s string, i, j int) (int, int) {
		l, r := i, j
		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			l--
			r++
		}
		return l + 1, r - 1
	}
	var res string
	for i := 0; i < len(s); i++ {
		l1, r1 := expand(s, i, i)
		if len(res) < r1-l1+1 {
			res = s[l1 : r1+1]
		}
		l2, r2 := expand(s, i, i+1)
		if len(res) < r2-l2+1 {
			res = s[l2 : r2+1]
		}
	}
	return res
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

func maxSubArray(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	var res = f[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(nums[i], f[i-1]+nums[i])
		res = max(res, f[i])
	}
	return res
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

func permute(nums []int) [][]int {
	var res [][]int
	var backTrack func(path []int, used uint64)
	backTrack = func(path []int, used uint64) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			masking := uint64(1 << i)
			if used&masking != 0 {
				continue
			}
			backTrack(append(path, nums[i]), used|masking)
		}
	}
	backTrack([]int{}, 0)
	return res
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		for j := len(nums) - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	l, r := i+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := (l + r) >> 1
	return mergeHerper(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeHerper(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeHerper(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeHerper(l1, l2.Next)
		return l2
	}
}

func lengthOfLIS(nums []int) int {
	var res []int
	for i := 0; i < len(nums); i++ {
		if len(res) == 0 || res[len(res)-1] < nums[i] {
			res = append(res, nums[i])
		} else {
			target := nums[i]
			idx := sort.Search(len(res), func(i int) bool {
				return res[i] >= target
			})
			if idx < len(res) {
				res[idx] = target
			}
		}
	}
	return len(res)
}

func isValid(s string) bool {
	var stack []byte
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if m[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func reorderList(head *ListNode) {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	newHead := reverseHelper(mid, nil)
	first, second := head, newHead
	for second != nil {
		firstNxt, secondNxt := first.Next, second.Next
		first.Next, second.Next = second, firstNxt
		first, second = firstNxt, secondNxt
	}
}

func minDistance(word1 string, word2 string) int {
	f := make([][]int, len(word1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(f); i++ {
		f[i][0] = i
	}
	for i := 0; i < len(f[0]); i++ {
		f[0][i] = i
	}
	for i := 0; i < len(word1)+1; i++ {
		for j := 0; j < len(word2)+1; j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1], f[i][j]) + 1
			}
		}
	}
	return f[len(word1)][len(word2)]
}

func addStrings(num1 string, num2 string) string {
	var (
		res   []byte
		carry int
	)
	i, j := len(num1)-1, len(num2)-1
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
		res = append([]byte{byte(sum + '0')}, res...)
	}
	if carry > 0 {
		res = append([]byte{byte(carry + '0')}, res...)
	}
	return string(res)
}

func maxPathSum(root *TreeNode) int {
	var res = math.MinInt
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(0, maxGain(root.Left))
		r := max(0, maxGain(root.Right))
		res = max(res, root.Val+l+r)
		//只有左子树或者右子树能构成路径
		return root.Val + max(l, r)
	}
	maxGain(root)
	return res
}

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

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
				res = max(res, f[i+1][j+1])
			}
		}
	}
	return res * res
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
	l := sortList(head)
	r := sortList(mid)
	return mergeHerper(l, r)
}

func longestValidParentheses(s string) int {
	//存储idx
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
func mySqrt(x int) int {
	idx := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if idx <= x && idx*idx == x {
		return idx
	}
	return idx - 1
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

func checkValidString(s string) bool {
	type elem struct {
		val byte
		idx int
	}
	s1, s2 := []elem{}, []elem{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			s1 = append(s1, elem{s[i], i})
		} else if s[i] == '*' {
			s2 = append(s2, elem{s[i], i})
		} else {
			if len(s1) > 0 {
				s1 = s1[:len(s1)-1]
			} else if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			} else {
				return false
			}
		}
	}
	if len(s1) == 0 {
		return true
	}
	i := len(s1) - 1
	j := len(s2) - 1
	for i >= 0 && j >= 0 {
		if s1[i].idx > s2[j].idx {
			return false
		}
		i--
		j--
	}
	return i < 0
}
