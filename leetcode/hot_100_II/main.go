package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

}

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for i := 0; i < len(strs); i++ {
		key := [26]int{}
		for _, ch := range strs[i] {
			key[ch-'a']++
		}
		m[key] = append(m[key], strs[i])
	}
	var res [][]string
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	var res int
	for i := 0; i < len(nums); i++ {
		if m[nums[i]-1] {
			continue
		}
		cnt := 0
		val := nums[i]
		for m[val] {
			val++
			cnt++
		}
		res = max(res, cnt)
	}
	return res
}

// 前缀和优化
func subarraySum(nums []int, k int) int {
	var res int
	var sum int
	prefix := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += prefix[sum-k]
		prefix[sum]++
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	prefix := map[int]int{0: 1}
	var backTrack func(root *TreeNode, sum int)
	backTrack = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		res += prefix[sum-targetSum]
		prefix[sum]++
		backTrack(root.Left, sum)
		backTrack(root.Right, sum)
		prefix[sum]--
	}
	backTrack(root, 0)
	return res
}

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var q []int //left side
	for r := 0; r < len(nums); r++ {
		//越界
		for len(q) > 0 && r-q[0]+1 > k {
			q = q[1:]
		}
		//单调递减队列
		for len(q) > 0 && nums[r] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		//0,1,2,3
		if r >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

// 原地hash或者map优化
func firstMissingPositive(nums []int) int {
	abs := func(val int) int {
		if val < 0 {
			return -val
		}
		return val
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -(abs(nums[num-1]))
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

// 转置+reverseHelper
func rotate(matrix [][]int) {
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
	for cur := head; cur.Next != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	newHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		newNode := cur.Next
		cur.Next = newNode.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}
	return newHead
}

// 二叉树
func diameterOfBinaryTree(root *TreeNode) int {
	var depth func(root *TreeNode) int
	var res int
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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			nxt := cur.Left
			pre := nxt
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, nxt
		}
		cur = cur.Right
	}
}

// 图DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	for _, p := range prerequisites {
		//学习p0，必须先学习p1
		from, to := p[0], p[1]
		graph[from] = append(graph[from], to)
	}
	visited := map[int]int{}
	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if visited[course] == 1 {
			return true
		}
		if visited[course] == 2 {
			return false
		}
		visited[course] = 1
		//dfs
		for _, nxt := range graph[course] {
			if hasCycle(nxt) {
				return true
			}
		}
		visited[course] = 2
		return false
	}
	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}
	return true
}

// 图BFS
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	var q []point
	var fresh int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, point{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	var res int
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(q) > 0 && fresh > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			cur := q[0]
			q = q[1:]
			for _, d := range directions {
				//nxt point
				ni := cur.x + d[0]
				nj := cur.y + d[1]
				if ni < 0 || ni > m-1 || nj < 0 || nj > n-1 || grid[ni][nj] != 1 {
					continue
				}
				grid[ni][nj] = 2
				fresh--
				q = append(q, point{ni, nj})
			}
		}
		res++
	}
	if fresh > 0 {
		return -1
	}
	return res
}

// top-k quick select优化
func findKthLargest(nums []int, k int) int {
	partition := func(nums []int, l, r int) int {
		random := l + rand.IntN(r-l+1)
		nums[r], nums[random] = nums[random], nums[r]
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
	var quickSelect func(nums []int, l, r int, topK int)
	quickSelect = func(nums []int, l, r int, topK int) {
		pivotIdx := partition(nums, l, r)
		if pivotIdx == topK {
			return
		} else if pivotIdx < topK {
			quickSelect(nums, pivotIdx+1, r, topK)
		} else {
			quickSelect(nums, l, pivotIdx-1, topK)
		}
	}
	quickSelect(nums, 0, len(nums)-1, len(nums)-k)
	return nums[len(nums)-k:][0]
}

// top-k heap优化
type minHeap struct {
	nums []int
}

func (h *minHeap) len() int {
	return len(h.nums)
}

func (h *minHeap) siftUp(i int) {
	parent := (i - 1) / 2
	if parent < 0 || h.nums[parent] <= h.nums[i] {
		return
	}
	h.nums[i], h.nums[parent] = h.nums[parent], h.nums[i]
	h.siftUp(parent)
}

// append->children want to up
func (h *minHeap) push(val int) {
	h.nums = append(h.nums, val)
	h.siftUp(len(h.nums) - 1)
}

func (h *minHeap) siftDown(i int) {
	l, r, smallest := 2*i+1, 2*i+2, i
	if l < len(h.nums) && h.nums[l] <= h.nums[smallest] {
		smallest = l
	}
	if r < len(h.nums) && h.nums[r] <= h.nums[smallest] {
		smallest = r
	}
	if smallest == i {
		return
	}
	h.nums[i], h.nums[smallest] = h.nums[smallest], h.nums[i]
	h.siftDown(smallest)
}

// swap to top,top wants to down
func (h *minHeap) pop() (int, bool) {
	if len(h.nums) == 0 {
		return -1, false
	}
	val := h.nums[0]
	h.nums[0], h.nums[len(h.nums)-1] = h.nums[len(h.nums)-1], h.nums[0]
	h.nums = h.nums[:len(h.nums)-1]
	h.siftDown(0)
	return val, true
}

func (h *minHeap) peek() int {
	return h.nums[0]
}

func findKthLargestHeap(nums []int, k int) int {
	h := &minHeap{}
	for i := 0; i < len(nums); i++ {
		if h.len() < k {
			h.push(nums[i])
		} else if h.nums[i] > h.peek() {
			h.pop()
			h.push(nums[i])
		}
		fmt.Println(h.peek())
	}
	return h.nums[0]
}

func partitionLabels(s string) []int {
	lastPos := map[byte]int{}
	for i := 0; i < len(s); i++ {
		lastPos[s[i]] = i
	}
	var res []int
	var start, end int
	for i := 0; i < len(s); i++ {
		end = max(end, lastPos[s[i]])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
