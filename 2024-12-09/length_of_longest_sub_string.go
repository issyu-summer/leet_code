package _024_12_09

import "math"

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	check := func(b byte) bool {
		return m[b] > 0
	}
	ans := math.MinInt
	l, r := 0, 0
	for ; r < len(s); r++ {
		for ; l <= r && check(s[l]); l++ {
			m[s[l]]--
		}
		if ans < r-l+1 {
			ans = r - l + 1
		}
		m[s[r]]++
	}
	return ans
}
