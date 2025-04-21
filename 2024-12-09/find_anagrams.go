package _024_12_09

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	anagrams := [26]int{}
	for _, ch := range p {
		anagrams[ch-'a']++
	}
	isAns := func(str string) bool {
		if len(str) < len(p) {
			return false
		}
		checkArr := [26]int{}
		for _, ch := range str {
			checkArr[ch-'a']++
		}
		return checkArr == anagrams
	}

	check := func(l, r int) bool {
		return r-l+1 > len(p)
	}

	for l, r := 0, 0; r < len(s); r++ {
		for ; l <= r && check(l, r); l++ {

		}
		if isAns(s[l : r+1]) {
			ans = append(ans, l)
		}
	}
	return ans
}

func findAnagrams1(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	ans := make([]int, 0)
	anagrams := [26]int{}
	checkArr := [26]int{}
	for i, _ := range p {
		anagrams[p[i]-'a']++
		checkArr[s[i]-'a']++
	}
	check := func(checkArr [26]int) bool {
		return checkArr == anagrams
	}
	if check(checkArr) {
		ans = append(ans, 0)
	}
	for i := len(p); i < len(s); i++ {
		//保持3位滑动
		checkArr[s[i]-'a']++
		checkArr[s[i-len(p)]-'a']--
		if check(checkArr) {
			ans = append(ans, i-len(p)+1)
		}
	}
	return ans
}
