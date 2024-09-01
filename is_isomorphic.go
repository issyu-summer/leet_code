package main

func main() {
	isIsomorphic("badc", "baba")
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
