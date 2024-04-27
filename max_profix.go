package main

import "fmt"

func main() {
	fmt.Print(maxProfit([]int{7, 1, 5, 3, 6}))
}

func maxProfit(prices []int) int {
	income := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > income {
			income = prices[i] - minPrice
		}
	}
	return income
}
