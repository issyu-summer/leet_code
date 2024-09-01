package main

import "strings"

func main() {
	wordPattern("aaaa", "dog cat cat dog")
}

func wordPattern(pattern string, s string) bool {
	pMap := make(map[byte]string)
	sMap := make(map[string]byte)
	strArr := strings.Split(s, " ")
	if len(pattern) != len(strArr) {
		return false
	}
	for i := 0; i < len(strArr); i++ {
		if v, ok := pMap[pattern[i]]; ok && v != strArr[i] {
			return false
		}
		if v, ok := sMap[strArr[i]]; ok && v != pattern[i] {
			return false
		}
		pMap[pattern[i]] = strArr[i]
		sMap[strArr[i]] = pattern[i]
	}
	return true
}
