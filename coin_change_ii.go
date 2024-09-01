package main

func main() {

}

func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] += f[j-coins[i]]
		}
	}
	return f[amount]
}
