package main

import "fmt"

func findRepeatNumber(documents []int) int {
	for i := 0; i < len(documents); i++ {
		if documents[documents[i]] != documents[i] {
			documents[i], documents[documents[i]] = documents[documents[i]], documents[i]
		}
	}
	fmt.Println(documents)
	return -1
}

func main() {
	//0,1,2,3
	ints := []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 120, 12, 13, 14, 15}
	findRepeatNumber(ints)
}
