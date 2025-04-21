package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	words := map[string]struct{}{}
	for _, word := range wordDict {
		words[word] = struct{}{}
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := words[s[j:i]]; ok && f[j] {
				f[i] = true
			}
		}
	}
	return f[len(s)]
}

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = n + 1
	}
	f[0] = 0
	f[1] = f[1-1*1] + 1 //拆分1的平方和0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := map[byte]int{}
	for i, j := 0, 0; j < len(s); j++ {
		if idx, ok := m[s[j]]; ok && idx >= i {
			i = idx + 1
		}
		m[s[j]] = j
		res = max(res, j-i+1)
	}
	return res
}

func lengthOfLongestSubstringII(s string) int {
	res := 0
	m := map[byte]int{}
	for i, j := 0, 0; j < len(s); j++ {
		for _, ok := m[s[j]]; ok && i < j; _, ok = m[s[j]] {
			delete(m, s[i])
			i++
		}
		m[s[j]] = j
		res = max(res, j-i+1)
	}
	return res
}
