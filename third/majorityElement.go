package main

func main() {

}

func majorityElement(nums []int) int {
	majority, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if majority == num {
			count++
		} else {
			count--
		}
	}
	return majority
}
