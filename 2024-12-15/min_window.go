package _024_12_15

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	sArr := [256]int{}
	tArr := [256]int{}
	for i := 0; i < len(t); i++ {
		tArr[t[i]]++
	}
	checkAns := func() bool {
		for i := 0; i < len(tArr); i++ {
			if tArr[i] == 0 {
				continue
			}
			if sArr[i] < tArr[i] {
				return false
			}
		}
		return true
	}
	last, ans := "", ""
	for l, r := 0, 0; r < len(s); r++ {
		sArr[s[r]]++
		for ; l <= r && checkAns(); l++ {
			sArr[s[l]]--
			last = s[l : r+1]
			if len(ans) > len(last) || ans == "" {
				ans = last
			}
		}
	}
	return ans
}
