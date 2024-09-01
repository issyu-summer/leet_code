package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		coinChangeDFS  func(coins []int, amount, count int)
		coinChangeMemo func(coins []int, amount int) int
		memo           = map[int]int{}
		res            = math.MaxInt32
	)
	coinChangeMemo = func(coins []int, amount int) int {
		//凑不成amount
		if amount < 0 {
			return -1
		}
		//能够凑成amount
		if amount == 0 {
			return 0
		}
		if memo[amount] != 0 {
			return memo[amount]
		}
		res := math.MaxInt32
		for _, coin := range coins {
			count := coinChangeMemo(coins, amount-coin)
			if count == -1 {
				continue
			}
			if count < res {
				res = count + 1
			}
		}
		if res == math.MaxInt32 {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
		return memo[amount]
	}
	coinChangeDFS = func(coins []int, amount, count int) {
		//凑不成amount
		if amount < 0 {
			return
		}
		//能够凑成amount
		if amount == 0 {
			res = min(res, count)
		}
		for _, c := range coins {
			coinChangeDFS(coins, amount-c, count+1)
		}
	}
	c := func(coins []int, amount int) int {
		if len(coins) == 0 {
			return -1
		}
		coinChangeDFS(coins, amount, 0)
		if res == math.MaxInt32 {
			return -1
		}
		return res
	}
	ans := c([]int{2}, 3)
	fmt.Println(ans)
}

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = amount + 1
	}
	f[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				f[i] = min(f[i], f[i-coins[j]]+1)
			}
		}
	}
	if f[amount] == amount+1 {
		return -1
	}
	return f[amount]
}

func coinChangeIII(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = amount + 1
	}
	f[0] = 0
	//外层循环物品
	for i := 0; i < len(coins); i++ {
		//内层循环容量，从小到大（完全背包）
		for j := coins[i]; j <= amount; j++ {
			f[j] = min(f[j], f[j-coins[i]]+1)
		}
	}
	if f[amount] == amount+1 {
		return -1
	}
	return f[amount]
}
