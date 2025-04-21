package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	cnt := map[rune]int{}
	for i := 0; i < len(s); i++ {
		cnt[rune(s[i])]++
		cnt[rune(t[i])]--

	}
	for _, val := range cnt {
		if val != 0 {
			return false
		}
	}
	return true
}
