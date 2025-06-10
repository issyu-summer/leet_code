package main

import "fmt"

func main() {
}

// 无重复的最长子串
func lengthOfLongestSubstring(s string) int {
	var res int
	exist := map[byte]bool{}
	for l, r := 0, 0; r < len(s); r++ {
		for l < r && exist[s[r]] {
			exist[s[l]] = false
			l++
		}
		//最长放后
		exist[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}

// 最小覆盖子串
func minWindow(s string, t string) string {
	var res = ""
	tM := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tM[t[i]]++
	}
	sM := map[byte]int{}
	check := func() bool {
		for key, tVal := range tM {
			if sVal, ok := sM[key]; !ok || (ok && sVal < tVal) {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < len(s); r++ {
		sM[s[r]]++
		for l < r && check() {
			if len(res) > r-l+1 || res == "" {
				res = s[l : r+1]
				fmt.Println(res)
			}
			sM[s[l]]--
			l++
		}
	}
	return res
}

// 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	var res = len(nums) + 1
	var sum int
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for l < r && sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
