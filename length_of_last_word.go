package main

func main() {

}

func lengthOfLastWord(s string) int {
	l := len(s) - 1
	for {
		if s[l] == ' ' {
			l--
		} else {
			break
		}
	}
	n := 0
	for i := l; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		n++
	}
	return n
}
