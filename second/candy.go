package main

func main() {

}

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for j := len(ratings) - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			candies[j] = max(candies[j], candies[j+1]+1)
		}
	}
	totalCandies := 0
	for i := 0; i < len(ratings); i++ {
		totalCandies += candies[i]
	}
	return totalCandies
}
