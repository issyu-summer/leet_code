package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitPhaseII(prices []int) int {
	income := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		income[i] = make([]int, 2)
	}
	income[0][0] = 0
	income[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		income[i][0] = max(income[i-1][0], income[i-1][1]+prices[i])
		income[i][1] = max(income[i-1][1], income[i-1][0]-prices[i])
	}
	return income[len(prices)-1][0]
}
