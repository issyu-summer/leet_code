package main

import "math"

func main() {

}

// 3,2,1,0,4
// 2,3,1,1,4
func canJump(nums []int) bool {
	maxPosition := 0
	for i := 0; i < len(nums); i++ {
		// 当前最远位置无法走到i，因为没办法继续走了
		if i > maxPosition {
			return false
		}
		// 如果能继续走则更新max position
		maxPosition = max(maxPosition, i+nums[i])
		// 如果最远位置>len-1,则一定能走到
		if maxPosition >= len(nums)-1 {
			return true
		}
	}
	return false
}

func jump(nums []int) int {
	cnt, end, maxPos := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if end == i {
			cnt++
			end = maxPos
		}
		if end >= len(nums)-1 {
			return cnt
		}
	}
	return 0
}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	res := 0
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		if res < prices[i]-minPrice {
			res = prices[i] - minPrice
		}
	}
	return res
}

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 0; i < len(f); i++ {
		f[i] = amount + 1
	}
	f[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			f[i] = min(f[i], f[i-coin]+1)
		}
	}
	if f[amount] == amount+1 {
		return -1
	}
	return f[amount]
}

func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			f[i] += f[i-coin]
		}
	}
	return f[amount]
}
