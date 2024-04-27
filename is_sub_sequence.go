package main

func main() {

}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i, j := 0, 0
	for i < len(t) {
		if t[i] == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
		i++
	}
	return false
}
