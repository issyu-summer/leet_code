package main

func main() {

}

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
	//1,1,1
	//1,2,3
	//1,3,6
	//=>
	//1,1,1
	//1,fj+=fj-1=2
	//if m > n {
	//	return uniquePaths(n, m)
	//}
	//f := make([]int, n)
	//for i := 0; i < len(f); i++ {
	//	f[i] = 1
	//}
	//for i := 1; i < m; i++ {
	//	for j := 1; j < n; j++ {
	//		f[j] += f[j-1]
	//	}
	//}
	//return f[n-1]
}

// 最小路径和
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	//f(i,j)为以text1中i结尾，text2中j结尾的最长公共子序列长度
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[m][n]
}

// 编辑距离
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < len(f); i++ {
		f[i][0] = i
	}
	for j := 0; j < len(f[0]); j++ {
		f[0][j] = j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j], f[i][j]) + 1
			}
		}
	}
	return f[m][n]
}

// 打家劫舍
func rob(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	f[1] = max(f[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f[i] = max(nums[i]+f[i-2], f[i-1])
	}
	return f[len(nums)-1]
}

// 完全平方数
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = n + 1
	}
	f[0] = 0
	f[1] = f[1-1*1] + 1
	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
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

// 单词拆分
func wordBreak(s string, wordDict []string) bool {
	exist := map[string]bool{}
	for i := 0; i < len(wordDict); i++ {
		exist[wordDict[i]] = true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if f[j] && exist[s[j:i]] {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
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
	target := sum / 2
	f := make([]bool, target+1)
	f[0] = true
	//11(1,10)=>10(1,9)===>1(1,0)true
	//11(5,6)=>10(5,5)====>6(5,1)true=>5(5,0)true
	//11(5,6)true
	for i := 1; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			f[j] = f[j] || f[j-nums[i]]
			if f[target] {
				return true
			}
		}
	}
	return false
}
