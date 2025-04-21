package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//isValidBST(&TreeNode{
	//	Val:  5,
	//	Left: &TreeNode{Val: 1},
	//	Right: &TreeNode{
	//		Val:   4,
	//		Left:  &TreeNode{Val: 2},
	//		Right: &TreeNode{Val: 6},
	//	},
	//})
	reverseMessage("a good   example")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int

	var inOrderTrace func(root *TreeNode)
	inOrderTrace = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTrace(root.Left)
		k--
		fmt.Println(root.Val)
		if k == 0 {
			res = root.Val
			return
		}
		inOrderTrace(root.Right)
	}
	inOrderTrace(root)
	return res
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
	return check(root, math.MinInt64, math.MaxInt64)
}

func removeDuplicates(nums []int) int {
	remove := func(nums []int, k int) int {
		if len(nums) < k {
			return len(nums)
		}
		ans := k
		for i, j := k, k; j < len(nums); j++ {
			if nums[i-k] != nums[j] {
				nums[i] = nums[j]
				i++
				ans = i
			}
		}
		return ans
	}
	return remove(nums, 2)
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func reverseMessage(message string) string {
	messages := strings.Fields(message)
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return strings.Join(messages, " ")
}

func reverseStr(message string) string {
	bytes := []byte(message)
	i, j := 0, len(message)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], message[i]
		i++
		j--
	}
	return string(bytes)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}

func lengthOfLIS(nums []int) int {
	var binarySearchGE func([]int, int) int
	ans := make([]int, 0)
	for _, num := range nums {
		if len(ans) == 0 || ans[len(ans)-1] < num {
			ans = append(ans, num)
		} else {
			idx := binarySearchGE(ans, num)
			ans[idx] = num
		}
	}
	return len(ans)
}
