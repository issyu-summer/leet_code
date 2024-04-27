package main

func main() {
	longestCommonPrefix([]string{"leets", "leetcode"})
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func lcp(strs1, strs2 string) string {
	var (
		l   = min(len(strs1), len(strs2))
		idx = 0
	)

	for i := 0; i < l; i++ {
		if strs1[i] == strs2[i] {
			idx++
		} else {
			break
		}
	}
	return strs1[:idx]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
