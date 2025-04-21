package main

import "math"

func main() {
	longestCommonPrefixII([]string{"flower", "flow", "flight"})
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

func lcp(str1, str2 string) string {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var (
		l   = min(len(str1), len(str2))
		idx = 0
	)
	for i := 0; i < l; i++ {
		if str1[i] == str2[i] {
			idx++
		} else {
			break
		}
	}
	return str1[:idx]
}

// 有公共特征
func longestCommonPrefixII(strs []string) string {
	min := func(strs ...string) int {
		ans := math.MaxInt32
		for _, v := range strs {
			if len(v) < ans {
				ans = len(v)
			}
		}
		return ans
	}
	isCommonPrefix := func(strs []string, length int) bool {
		str0 := strs[0][:length]
		for _, str := range strs[1:] {
			if str[:length] != str0 {
				return false
			}
		}
		return true
	}
	//都是取到上边界
	i, j := 0, min(strs...)
	for i < j {
		//上中位数
		mid := i + (j-i+1)>>1
		if isCommonPrefix(strs, mid) {
			i = mid
		} else {
			j = mid - 1
		}
	}
	return strs[0][:i]
}

func longestCommonPrefixIII(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
