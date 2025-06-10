package main

func main() {
}

// code tap*4
// 1.合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, end := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[end] = nums1[i]
			end--
			i--
		} else {
			nums1[end] = nums2[j]
			end--
			j--
		}
	}
	for i >= 0 {
		nums1[end] = nums1[i]
		end--
		i--
	}
	for j >= 0 {
		nums1[end] = nums2[j]
		end--
		j--
	}
}

// 2.买卖股票的最佳时机
func maxProfit(prices []int) int {
	res := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		res = max(res, prices[i]-minPrice)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 3.二叉树的Z形便利
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	leftToRight := true
	for len(q) > 0 {
		levelSize := len(q)
		var levelRes []int
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if leftToRight {
				levelRes = append(levelRes, node.Val)
			} else {
				levelRes = append(append([]int{}, node.Val), levelRes...)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, levelRes)
		leftToRight = !leftToRight
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 4.链表中是否有环
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
