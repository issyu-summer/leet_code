package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	var ar []string
	for _, t := range split {
		if t != "" {
			ar = append(ar, t)
		}
	}
	for i, _ := range ar {
		ar[i] = reverseStr(ar[i])
	}
	return reverseStr(strings.Join(ar, " "))
}

func reverseStr(s string) string {
	bytes := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
