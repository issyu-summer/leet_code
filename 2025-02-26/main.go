package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	f := make([]int, len(prices))
	minCost := math.MaxInt64
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < minCost {
			minCost = prices[i-1]
		}
		f[i] = max(f[i-1], prices[i]-minCost)
	}
	return f[len(prices)-1]
}

func maxProfit1(prices []int) int {
	f := 0
	minCost := math.MaxInt64
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < minCost {
			minCost = prices[i-1]
		}
		f = max(f, prices[i]-minCost)
	}
	return f
}
