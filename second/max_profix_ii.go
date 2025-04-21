package main

func main() {

}

func maxProfitPhaseII(prices []int) int {
	return dp(prices)
}

func greedy(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func dp(prices []int) int {
	f := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		f[i] = make([]int, 2)
	}
	f[0][0] = 0
	f[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1]+prices[i])
		f[i][1] = max(f[i-1][1], f[i-1][0]-prices[i])
	}
	//最后一天一定卖出，利润最大
	return f[len(prices)-1][0]
}
