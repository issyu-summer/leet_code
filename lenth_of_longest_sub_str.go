package main

func main() {
}

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	ans := 0
	for l, r := 0, 0; r < len(s); r++ {
		for l <= r && set[s[r]] {
			set[s[l]] = false
			l++
		}
		set[s[r]] = true
		if ans < r-l+1 {
			ans = r - l + 1
		}
	}
	return ans
}
