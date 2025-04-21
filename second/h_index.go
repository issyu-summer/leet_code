package main

func main() {

}

func hIndex(citations []int) int {
	cnt := make([]int, len(citations)+1)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations) {
			cnt[len(citations)]++
		} else {
			cnt[citations[i]]++
		}
	}
	for i, tot := len(citations), 0; i >= 0; i-- {
		tot += cnt[i]
		if tot >= i {
			return i
		}
	}
	return 0
}
