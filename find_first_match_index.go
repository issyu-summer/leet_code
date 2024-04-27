package main

func main() {

}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		s1, s2 := i, 0
		for s1 < len(haystack) && s2 < len(needle) && haystack[s1] == needle[s2] {
			s1++
			s2++
		}
		if s2 == len(needle) {
			return i
		}
	}
	return -1
}
