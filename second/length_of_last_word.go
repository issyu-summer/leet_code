package main

func main() {

}

func lengthOfLastWord(s string) int {
	l := len(s) - 1
	ans := 0
	for ; l >= 0; l-- {
		if s[l] == ' ' {
			continue
		} else {
			break
		}
	}
	for ; l >= 0 && s[l] != ' '; l-- {
		ans++
	}
	return ans
}
