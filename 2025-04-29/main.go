package main

import (
	"sort"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	//使用dummy简化头节点的删除操作
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			val := cur.Val
			for cur.Val == val {
				cur = cur.Next
			}
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}

// 左边一半，一定下于右边一半
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		findMedianSortedArrays(nums1, nums2)
	}
	m, n := len(nums1), len(nums2)
	total := m + n
	//取多一半,左边多，右边少
	half := (total + 1) / 2
	if m == 0 {
		if total%2 == 1 {
			return float64(nums2[half-1])
		}
		return float64(nums2[half-1]+nums2[half]) / 2.0
	}
	i := sort.Search(m, func(i int) bool {
		//nums1左边为i,nums2左边为j
		j := half - i
		//j越界，且j过大，需要增大i，即向右搜索
		if j > n {
			return false
		}
		//j越界，且j过小，需要减小i，即向左搜索
		if j < 0 {
			return true
		}
		//至此，0<=j&&j<=n 且0<=i&&i<=m

		//i,j是mid,或者可以理解为右最小
		//i-1,j-1是左最大
		//nums2 left mid > nums1 right mid,不符合规则,j过大
		if j > 0 && i < m && nums2[j-1] > nums1[i] {
			return false
		}
		//nums1 left mid > nums2 right mid，不符合规则，i过大
		if i > 0 && j < n && nums1[i-1] > nums2[j] {
			return true
		}
		//如果都符合，继续向左搜索，找到最小的
		return true
	})
	//j-1 left max,j right min
	j := half - i
	var maxLeft int
	if i == 0 {
		maxLeft = nums2[j-1]
	} else if j == 0 {
		maxLeft = nums1[i-1]
	} else {
		maxLeft = max(nums1[i-1], nums2[j-1])
	}
	if total%2 == 1 {
		return float64(maxLeft)
	}
	var minRight int
	if i == m {
		minRight = nums2[j]
	} else if j == n {
		minRight = nums1[i]
	} else {
		minRight = min(nums1[i], nums2[j])
	}
	return float64(maxLeft+minRight) / 2.0
}
