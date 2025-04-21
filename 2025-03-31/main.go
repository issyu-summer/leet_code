package main

import (
	"slices"
	"sort"
)

func main() {
	hIndex([]int{3, 0, 6, 1, 5})
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func change(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func maxProfit(prices []int) int {
	maxP := func(prices []int) int {
		res := 0
		for i := 1; i < len(prices); i++ {
			res += max(0, prices[i]-prices[i-1])
		}
		return res
	}
	return maxP(prices)
	//dp := make([][2]int, len(prices))
	//dp[0][1] = -prices[0]
	//for i := 1; i < len(prices); i++ {
	//	dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
	//	dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	//}
	//return dp[len(dp)-1][0]
}

func jump(nums []int) int {
	cnt, curEnd, maxPosition := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		//当i达到cur end，则不得不往前跳一步
		if i == curEnd {
			cnt++
			curEnd = maxPosition
			if curEnd >= len(nums)-1 {
				break
			}
		}
	}
	return cnt
}

func canJump(nums []int) bool {
	curEnd, maxPosition := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == curEnd {
			curEnd = maxPosition
		}
	}
	return curEnd >= len(nums)-1
}

func hIndex(citations []int) int {
	slices.SortFunc(citations, func(a, b int) int {
		return b - a
	})
	//6,5,3,1,0
	var h int
	for i := 0; i < len(citations); i++ {
		if citations[i] > h {
			h++
		}
	}
	return h
}

func hIndexII(citations []int) int {
	n := len(citations)
	//
	return n - sort.Search(n, func(x int) bool {
		return citations[x] >= n-x
	})
}
