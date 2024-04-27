package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	memo := map[int]bool{}
	return canBreak(0, s, wordDict, memo)
}

func InDict(target string, wordDict []string) bool {
	for _, s := range wordDict {
		if target == s {
			return true
		}
	}
	return false
}

func canBreak(start int, s string, wordDict []string, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	if res, ok := memo[start]; ok {
		return res
	}
	//左闭右开
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if InDict(prefix, wordDict) && canBreak(i, s, wordDict, memo) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}
