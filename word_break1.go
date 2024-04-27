package main

func main() {

}

func wordBreak1(s string, wordDict []string) bool {
	wordM := make(map[string]bool, 0)
	for _, w := range wordDict {
		wordM[w] = true
	}
	memo := map[int]bool{}
	return canBreak1(0, s, wordM, memo)
}

func canBreak1(start int, s string, wordMap map[string]bool, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	if res, ok := memo[start]; ok {
		return res
	}
	//左闭右开
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreak1(i, s, wordMap, memo) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}
