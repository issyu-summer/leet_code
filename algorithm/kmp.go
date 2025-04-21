package main

import "fmt"

func main() {
	text := "ABABDABACDABABCABB"
	pattern := "ABABCABAB"
	idx := KMPSearch(text, pattern)
	fmt.Println(idx)
}

// lps定义，对于前缀和后缀，其中最长的相等的前缀和后缀的长度
func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length, i := 0, 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func KMPSearch(text, pattern string) []int {
	ans := make([]int, 0)
	lps := buildLPS(pattern)
	i, j := 0, 0
	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}
		if j == len(pattern) {
			//fmt.Printf("Pattern found at index %d\n", i-j)
			ans = append(ans, i-j)
			j = lps[j-1]
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return ans
}
