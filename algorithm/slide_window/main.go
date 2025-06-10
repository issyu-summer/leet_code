package main

func main() {

}

func minSubArrayLen(target int, nums []int) int {
	var res = len(nums) + 1
	var sum int
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		//check r的影响
		for l <= r && sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if res == len(nums)+1 {
		return -1
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	var res = 0
	m := map[byte]struct{}{}
	//check r的影响
	for l, r := 0, 0; r < len(s); r++ {
		for _, ok := m[s[r]]; l <= r && ok; _, ok = m[s[r]] {
			delete(m, s[l])
			l++
		}
		m[s[r]] = struct{}{}
		res = max(res, r-l+1)
	}
	return res
}

func minWindow(s string, t string) string {
	sMap := map[byte]int{}
	tMap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	check := func() bool {
		for key, val := range tMap {
			if val > sMap[key] {
				return false
			}
		}
		return true
	}
	var res = s + "#"
	for l, r := 0, 0; r < len(s); r++ {
		sMap[s[r]]++
		for l <= r && check() {
			if len(res) > len(s[l:r+1]) {
				res = s[l : r+1]
			}
			sMap[s[l]]--
			l++
		}
	}
	if res == s+"#" {
		return ""
	}
	return res
}
