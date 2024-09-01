package main

import "fmt"

func main() {
	fmt.Print(minWindow("a", "a"))

}

func minWindow(s string, t string) string {
	ansL, ansR := -1, -1
	result := len(s) + 1
	targetM, countM := make(map[byte]int), make(map[byte]int)
	for _, b := range []byte(t) {
		targetM[b]++
	}
	check := func() bool {
		for b, v := range targetM {
			if countM[b] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < len(s); r++ {
		countM[s[r]]++
		for l <= r && check() {
			if result > r-l+1 {
				result = r - l + 1
				ansL = l
				ansR = r
			}
			countM[s[l]]--
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL : ansR+1]
}
