package main

func main() {

}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n)
	}
	var maxSide int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := matrix[i][j] - '0'
			if num == 0 {
				continue
			}
			if i == 0 || j == 0 {
				f[i][j] = 1
			} else {
				f[i][j] = min(f[i][j-1], f[i-1][j], f[i-1][j-1]) + 1
			}
			maxSide = max(f[i][j], maxSide)
		}
	}
	return maxSide * maxSide
}

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i, j := m-1, 0; i >= 0 && j <= n-1; {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	mid := func(head *ListNode) *ListNode {
		//dummy(s,f)->1(s)->2(f)(2)->3->nil(f)
		//dummy(s,f)->1(s)->2(f)(s)->3->4(f)->nil => s=2 成立
		//dummy := &ListNode{Next: head}
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}
	m := mid(head)
	reverse := func(start, end *ListNode) *ListNode {
		var pre, cur *ListNode = nil, start
		for cur != end {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		return pre
	}
	mNxtH := reverse(m, nil)
	cur1, cur2 := head, mNxtH
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return true
}
