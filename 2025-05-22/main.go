package main

import "fmt"

func main() {

}

// 长度最小，考虑DP
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) < len(nums2) {
		return findLength(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	f := make([]int, m+1)
	var res int
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				f[j] = f[j-1] + 1
			} else {
				f[j] = 0
			}
			res = max(res, f[j])
		}
	}
	return res
}

func wordBreak(s string, wordDict []string) bool {
	f := make([]bool, len(s)+1)
	f[0] = true
	dict := map[string]bool{}
	for _, word := range wordDict {
		dict[word] = true
	}
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if f[j] && dict[s[j:i]] {
				f[i] = true
			}
		}
	}
	return f[len(s)]
}

func minSubArrayLen(target int, nums []int) int {
	var res = len(nums) + 1
	var sum int
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		//l<=r，代表窗口大小至少为1,l<r代表窗口大小至少为2(r-l+1为窗口大小)
		for l <= r && sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
			fmt.Println(l, r)
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
