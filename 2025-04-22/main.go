package main

import (
	"container/heap"
	"sort"
)

func main() {

}

//code top *4

// 1.搜索排序旋转数组
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	idx := sort.Search(len(nums), func(i int) bool {
		//在同一段
		if (nums[i] >= nums[0]) == (target >= nums[0]) {
			return nums[i] >= target
		}
		//不在同一段
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

// 2.岛屿数量
func numIslands(grid [][]byte) int {
	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
	var res int
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

// 3.无重复全排列，回溯
func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int, uint64)
	backtrack = func(path []int, used uint64) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			masking := uint64(1 << i)
			if used&masking != 0 {
				continue
			}
			backtrack(append(path, nums[i]), used|masking)
		}
	}
	backtrack([]int{}, 0)
	return res
}

// 4.有效的括号
func isValid(s string) bool {
	//遇到)要找到对应的(出列
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && m[stack[len(stack)-1]] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// leetcode hot 100
type Element struct {
	value int
	freq  int
}

type MinHeap []Element

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].freq < m[j].freq
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Element))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

// 1.前k个高频元素
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}
	h := &MinHeap{}
	heap.Init(h)
	for num, cnt := range freq {
		if h.Len() < k {
			heap.Push(h, Element{num, cnt})
		} else if cnt > ((*h)[0].freq) {
			heap.Pop(h)
			heap.Push(h, Element{num, cnt})
		}
	}
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(Element).value)
	}
	return res
}

// leetcode 150
type Elem struct {
	x    int
	y    int
	xIdx int
	yIdx int
}

type ElemHeap []Elem

func (e ElemHeap) Len() int {
	return len(e)
}

func (e ElemHeap) Less(i, j int) bool {
	return e[i].x+e[i].y < e[j].x+e[j].y
}

func (e ElemHeap) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *ElemHeap) Push(x any) {
	*e = append(*e, x.(Elem))
}

func (e *ElemHeap) Pop() any {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &ElemHeap{}
	heap.Init(h)
	//遍历全部组合，超过时间限制
	//for i := 0; i < len(nums1); i++ {
	//	for j := 0; j < len(nums2); j++ {
	//		if h.Len() < k {
	//			heap.Push(h, Elem{nums1[i], nums2[j]})
	//		} else if (*h)[0].x+(*h)[0].y > nums1[i]+nums2[j] {
	//			heap.Pop(h)
	//			heap.Push(h, Elem{nums1[i], nums2[j]})
	//		}
	//	}
	//}
	for i := 0; i < k && i < len(nums1); i++ {
		heap.Push(h, Elem{nums1[i], nums2[0], i, 0})
	}
	var res [][]int
	for h.Len() > 0 && len(res) < k {
		elem := heap.Pop(h).(Elem)
		x, y := elem.x, elem.y
		xIdx, yIdx := elem.xIdx, elem.yIdx
		res = append(res, []int{x, y})
		if yIdx+1 < len(nums2) {
			heap.Push(h, Elem{nums1[xIdx], nums2[yIdx+1], xIdx, yIdx + 1})
		}
	}
	return res
}
