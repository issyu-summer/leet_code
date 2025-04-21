package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"sort"
)

func main() {

}

//codetop

// 1.滑动窗口
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	res := 0
	set := map[byte]bool{}
	for i, j := 0, 0; j < len(s); j++ {
		for i < j && set[s[j]] {
			delete(set, s[i])
			i++
		}
		set[s[j]] = true
		if res < j-i+1 {
			start, end = i, j
			res = j - i + 1
		}
	}
	fmt.Println("res", s[start:end+1])
	return res
}

// LRUCache 1.常用算法-LRU
type LRUCache struct {
	cache    map[int]*list.Element
	list     *list.List
	capacity int
}
type KeyValue struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     list.New(),
		cache:    make(map[int]*list.Element),
		capacity: capacity,
	}
}

// Get 获取并移动到队头
func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(*KeyValue).val
	}
	return -1
}

// Put 存在则更新元素并移动到头、不存在则插入元素，并移动到头(如果长度超限，则淘汰元素)
func (lru *LRUCache) Put(key int, value int) {
	// update elem & move to front
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		elem.Value = &KeyValue{key, value}
	} else {
		//evict tail elem
		if lru.list.Len() >= lru.capacity {
			tail := lru.list.Back()
			lru.list.Remove(tail)
			delete(lru.cache, tail.Value.(*KeyValue).key)
		}
		//insert elem at front
		newElem := lru.list.PushFront(&KeyValue{key, value})
		lru.cache[key] = newElem
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 3.反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// IntHeap 优先队列
type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 4. 数组中的第k大元素
func findKthLargest(nums []int, k int) int {
	h := new(IntHeap)
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func findKthLargestQuickSort(nums []int, k int) int {
	partition := func(nums []int, l, r int) int {
		pivot := nums[r]
		i := l
		for j := l; j < r; j++ {
			if nums[j] <= pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[r] = nums[r], nums[i]
		return i
	}
	target := len(nums) - k
	l, r := 0, len(nums)-1
	for {
		idx := partition(nums, l, r)
		if idx == target {
			return nums[idx]
		} else if idx < target {
			l = idx + 1
		} else {
			r = idx - 1
		}
	}
}

// 经典起150
// 1.合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	tail := len(nums1) - i
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[tail] = nums1[i]
			i--
		} else {
			nums1[tail] = nums2[j]
			j--
		}
		tail--
	}
	for i >= 0 {
		nums1[tail] = nums1[i]
		tail--
		i--
	}
	for j >= 0 {
		nums1[tail] = nums2[j]
		tail--
		j--
	}
}

// 2.原地移除元素
func removeElement(arr []int, val int) int {
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] == val {
			arr[i] = arr[j]
			j--
		} else {
			i++
		}
	}
	return i
}

// 3.删除有序数组中的重复元素
func removeDuplicates(arr []int) int {
	i, j := 1, 1
	for j < len(arr) {
		if arr[j] != arr[i-1] {
			arr[i] = arr[j]
			i++
		}
		j++
	}
	return i
}

// 4.删除有序数组中的重复元素
func removeDuplicatesII(nums []int) int {
	remove := func(nums []int, k int) int {
		if len(nums) <= k {
			return len(nums)
		}
		i, j := k, k
		for j < len(nums) {
			if nums[i-k] != nums[j] {
				nums[i] = nums[j]
				i++
			}
			j++
		}
		return i
	}
	return remove(nums, 2)
}

// 5.多数元素
func majorityElement(nums []int) int {
	res := 0
	cnt := 0
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

// 6.轮转数组
func rotate(nums []int, k int) {
	reverse := func(nums []int) {
		i, j := 0, len(nums)-1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

//leetcode hot 100

// 1.Trie 前缀树
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func ConstructorTrie() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// 2.课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	//from->[]to
	graph := make(map[int][]int)
	for _, p := range prerequisites {
		from, to := p[0], p[1]
		graph[from] = append(graph[from], to)
	}
	visited := map[int]int{}
	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if visited[course] == 1 {
			return true
		}
		if visited[course] == 2 {
			return false
		}
		visited[course] = 1
		for _, next := range graph[course] {
			if hasCycle(next) {
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

// 3.腐烂的橘子
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	//右，下，左，下
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	type point struct{ x, y int }
	var queue []point
	fresh := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, point{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	res := 0
	for len(queue) > 0 && fresh > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, d := range directions {
				ni := cur.x + d[0]
				nj := cur.y + d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					fresh--
					queue = append(queue, point{ni, nj})
				}
			}
		}
		res++
	}
	if fresh > 0 {
		return -1
	}
	return res
}

// 4.寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	total := m + n
	half := (total + 1) / 2
	//处理m为0的特殊情况
	if m == 0 {
		if total%2 == 1 {
			return float64(nums2[half-1])
		}
		return float64(nums2[half-1]+nums2[half]) / 2.0
	}
	//核心在于找到i,j满足条件：左半务必小于右半，包括：
	//nums1 left mid < nums1 right mid(nums1有序，天然成立)
	//nums2 left mid < nums2 right mid(nums2有序，天然成立)
	//nums1 left mid < nums2 right mid，不成立需要根据数组有序的特征=>调整i和j的大小
	//nums2 left mid < nums1 right mid，不成立需要根据数组有序的特征=>调整i和j的大小
	i := sort.Search(m, func(i int) bool {
		//2个数组合并后，确保左半部分的元素数量为half，即nums1有i个，nums2有j个，i+j=half
		j := half - i
		//如果j>n越界，则说明i太小，所以需要向右搜索
		if j > n {
			return false
		}
		//如果j<0越界，则说明i太大，所以应该向左搜索
		if j < 0 {
			return true
		}
		//至此，因为没有return，此时j>=0 && j<=n
		//且根据func定义，此时，i>=0 && i<=m

		//nums2的左半，比nums1右半大，因此j是错误，具体的，i过小，j过大，所以需要增大i，减小j
		if j > 0 && i < m && nums2[j-1] > nums1[i] {
			return false //增大i
		}
		//nums1的左半，比nums2右半大，因此i是错误的，具体的，i过大，j过小，所以需要减小i，增大j
		if i > 0 && j < n && nums1[i-1] > nums2[j] {
			return true
		}
		// 都符合条件，因此继续向左边搜索
		return true
	})

	j := half - i
	//i,j的含义,第一个符合：是nums2的左半，小于nums1右半；且nums1的左半，小于nums2的右半的i和j的位置
	var maxLeft int
	//i为0，则整个nums1都是右半，=> nums2[:j]=>nums1
	if i == 0 {
		maxLeft = nums2[j-1]
		//j=0,则整个nums2都是右半=>nums[:i]=>nums2
	} else if j == 0 {
		maxLeft = nums1[i-1]
	} else {
		//如果都不为0，则nums[:i]=>nums[j:] 且nums[:j]=>nums[:i]，即maxLeft可能是二者中的最小值
		maxLeft = max(nums1[i-1], nums2[j-1])
	}
	//若是奇数个，直接返回maxLeft
	if total%2 != 0 {
		return float64(maxLeft)
	}
	var minRight int
	//i为m，则整个nums1都是左半，=>nums1=>nums2[j:]
	if i == m {
		minRight = nums2[j]
		//j为n，则整个nums2都是左半，=>nums2=>nums1[i:]
	} else if j == n {
		minRight = nums1[i]
	} else {
		//如果都不为0，则nums[:i]=>nums[j:] 且nums[:j]=>nums[:i]，即maxLeft可能是二者中的最小值
		minRight = min(nums1[i], nums2[j])
	}
	return float64(maxLeft+minRight) / 2.0
}
