package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := [26]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		sMap[t[i]-'a']--
	}
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}
	return true
}
