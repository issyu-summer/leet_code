package main

func main() {

}

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	right := make([]int, len(ratings))
	for i := 0; i < len(left); i++ {
		left[i] = 1
		right[i] = 1
	}
	for i := 1; i < len(ratings)-1; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	count := 0
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			count += left[i]
		} else {
			count += right[i]
		}
	}
	return count
}
