package main

import "strings"

func main() {
	wordPattern("aaaa", "dog cat cat dog")
}

func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}
	pMap := make(map[byte]string)
	sMap := make(map[string]byte)
	for i, str := range arr {
		if targetStr, ok := pMap[pattern[i]]; ok && str != targetStr {
			return false
		}
		if targetByte, ok := sMap[str]; ok && targetByte != pattern[i] {
			return false
		}
		pMap[pattern[i]] = str
		sMap[str] = pattern[i]
	}
	return true
}
