package main

func main() {

}

func climbStairs(n int) int {
	return climbMemo(n, make([]int, n+1))
}

func climbDp(n int) int {
	dp := make([]int, n+1)
	//已经在记忆了
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbMemo(n int, memo []int) int {
	if n == 1 || n == 2 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = climbMemo(n-1, memo) + climbMemo(n-2, memo)
	return memo[n]
}
