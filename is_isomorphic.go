package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sArr := [128]byte{}
	tArr := [128]byte{}
	for i := 0; i < len(s); i++ {
		if sArr[s[i]] != 0 && sArr[s[i]] != t[i] {
			return false
		}
		if tArr[t[i]] != 0 && tArr[t[i]] != s[i] {
			return false
		}
		sArr[s[i]] = t[i]
		tArr[t[i]] = s[i]
	}
	return true
}

func isIsomorphic1(s string, t string) bool {
	sArr := map[byte]byte{}
	tArr := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if b, ok := sArr[s[i]]; ok && b != t[i] {
			return false
		}
		if b, ok := tArr[t[i]]; ok && b != s[i] {
			return false
		}
		sArr[s[i]] = t[i]
		tArr[t[i]] = s[i]
	}
	return true
}
