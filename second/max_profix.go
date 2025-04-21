package main

import "math"

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}

func maxProfit(prices []int) int {
	f := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		f[i] = max(f[i-1], prices[i]-min(prices[:i]...))
	}
	return f[len(prices)-1]
}

func maxProfitII(prices []int) int {
	//记忆化
	var (
		minPrice = math.MaxInt32
		f        = make([]int, len(prices))
	)

	// 外层循环物品
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i-1])
		f[i] = max(f[i-1], prices[i]-minPrice)
	}
	return f[len(prices)-1]
}

func min(arr ...int) int {
	ans := arr[0]
	for _, val := range arr {
		if val < ans {
			ans = val
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
