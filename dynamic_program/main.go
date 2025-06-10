package main

func main() {

}

// DP-极值---------------------------------------------------------------
// 最长回文子串
func longestPalindrome(s string) string {
	f := make([][]bool, len(s))
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, len(f))
	}
	//handle 1 char
	for i := 0; i < len(s); i++ {
		f[i][i] = true
	}
	//handle 2 char
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			f[i][i+1] = true
		}
	}
	//handle 3 char,逆序，避免未初始化
	for i := len(s) - 3; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	var res string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if f[i][j] && len(s) < len(s[i:j+1]) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

// 最大正方形
func maximalSquare(matrix [][]byte) int {
	f := make([][]int, len(matrix))
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(matrix[0]))
	}
	var maxSide int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				f[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			}
			maxSide = max(maxSide, f[i][j])
		}
	}
	return maxSide * maxSide
}

// 打家劫舍
func rob(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	if len(nums) == 1 {
		return nums[0]
	}
	f[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f[i] = max(f[i-1], nums[i]+f[i-2])
	}
	return f[len(nums)-1]
}

// 最大子数组和
func maxSubArray(nums []int) int {
	var res = nums[0]
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(nums[i], f[i-1]+nums[i])
		res = max(res, f[i])
	}
	return res
}

// 最小路径和
func minPathSum(grid [][]int) int {
	f := make([][]int, len(grid))
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				f[i][j] = grid[0][0]
			} else if i == 0 {
				f[i][j] = f[i][j-1] + grid[i][j]
			} else if j == 0 {
				f[i][j] = f[i-1][j] + grid[i][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
			}
		}
	}
	return f[len(grid)-1][len(grid[0])-1]
}

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	f := make([][]int, len(text1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return f[len(text1)][len(text2)]
}

// 乘积最大的子数组
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	minF, maxF := make([]int, len(nums)), make([]int, len(nums))
	minF[0], maxF[0] = nums[0], nums[0]
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			minF[i-1], maxF[i-1] = maxF[i-1], minF[i-1]
		}
		maxF[i] = max(nums[i], maxF[i-1]*nums[i])
		minF[i] = min(nums[i], minF[i-1]*nums[i])
		res = max(res, maxF[i])
	}
	return res
}

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	f := make([]int, len(nums))
	for i := 0; i < len(f); i++ {
		f[i] = 1
	}
	var res int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		res = max(res, f[i])
	}
	return res
}

// 零钱兑换
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 0; i < len(f); i++ {
		f[i] = amount + 1
	}
	f[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] = min(f[j], f[j-coins[i]]+1)
		}
	}
	if f[amount] == amount+1 {
		return -1
	}
	return f[amount]
}

// 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	f := make([][]int, len(nums1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(nums2)+1)
	}
	var res int
	for i := 1; i < len(nums1)+1; i++ {
		for j := 1; j < len(nums2)+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
				res = max(res, f[i][j])
			}
		}
	}
	return res
}

// DP-其他---------------------------------------------------------------
// 不同路径
func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		f[i][0] = 1
	}
	for j := 0; j < n; j++ {
		f[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}

// 爬楼梯
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	f[2] = 2
	for i := 3; i < n+1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 编辑距离
func minDistance(word1 string, word2 string) int {
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}
	f := make([][]int, len(word1))
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(word2))
	}
	for i := 0; i < len(word1); i++ {
		f[i][0] = i
	}
	for j := 0; j < len(word2); j++ {
		f[0][j] = j
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			}
		}
	}
	return f[len(word1)-1][len(word2)-1]
}

// 杨辉三角形
func generate(numRows int) [][]int {
	f := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		f[i] = make([]int, i+1)
		f[i][0] = 1
		f[i][i] = 1
		for j := 1; j < i; j++ {
			f[i][j] = f[i-1][j-1] + f[i-1][j]
		}
	}
	return f
}

// 单词拆分
func wordBreak(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if f[j] && dict[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

// 接雨水
func trap(height []int) int {
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	var res int
	for i := 0; i < len(height); i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

// 完全平方数
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = n + 1
	}
	f[1] = 1
	for i := 1; i < n; i++ {
		for j := 1; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

// 分割等和子集
func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	//转换为是否nums[...]的和是否为target
	target := sum / 2
	f := make([]bool, target+1)
	f[0] = true
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			return false
		}
		for j := target; j >= nums[i]; j-- {
			f[j] = f[j] || f[j-nums[i]]
		}
	}
	return f[target]
}

// 零钱兑换II
func coinChangeII(coins []int, amount int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[i] += f[i-coins[i]]
		}
	}
	return f[amount]
}
