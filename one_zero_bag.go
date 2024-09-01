package main

func main() {

}

func ZeroOneBag(weights []int, values []int, bagWeight int) int {
	dp := make([][]int, len(weights)+1)
	for i := range dp {
		dp[i] = make([]int, bagWeight)
	}
	for i := 1; i < len(weights)+1; i++ {
		for j := 1; j <= bagWeight; j++ {
			if j >= weights[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(weights)][bagWeight]
}
