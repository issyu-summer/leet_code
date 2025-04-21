package _024_12_02

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, ch := range str {
			cnt[ch-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}
	result := [][]string{}
	for _, strArr := range m {
		result = append(result, strArr)
	}
	return result
}

func groupAnagrams2(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		m[string(bytes)] = append(m[string(bytes)], str)
	}
	result := [][]string{}
	for _, strArr := range m {
		result = append(result, strArr)
	}
	return result
}
