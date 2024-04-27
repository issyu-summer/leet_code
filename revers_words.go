package main

import "strings"

func main() {
	reverseWords("the sky is blue")
}

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	ar := make([]string, 0)
	for _, s2 := range split {
		if s2 != "" {
			ar = append(ar, s2)
		}
	}
	s = strings.Join(ar, " ")
	s = revers(s, 0, len(s)-1)
	i, j := 0, 0
	for j < len(s) {
		if s[j] == ' ' {
			s = revers(s, i, j-1)
			i = j + 1
		}
		j++
	}
	s = revers(s, i, len(s)-1)
	return s
}

func revers(s string, l, r int) string {
	bytes := []byte(s)
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	return string(bytes)
}
