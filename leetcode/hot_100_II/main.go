package main

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
