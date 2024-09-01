package main

func main() {

}

func completeKnapsack(weights, values []int, capacity int) int {
	f := make([]int, capacity+1)
	for i := 0; i < len(weights); i++ {
		for j := weights[i]; j <= capacity; j++ {
			f[j] = max(f[j], f[j-weights[i]]+values[i])
		}
	}
	return f[capacity]
}
