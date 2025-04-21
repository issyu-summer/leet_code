package main

func main() {

}

func partition(s string) [][]string {
	isPalindrome := func(s string, i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var res [][]string
	var fc func(int, []string)
	fc = func(start int, path []string) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			if isPalindrome(s, start, i) {
				fc(i+1, append(path, s[start:i+1]))
			}
		}
	}
	fc(0, []string{})
	return res
}
