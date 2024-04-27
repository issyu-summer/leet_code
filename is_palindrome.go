package main

import (
	"fmt"
	"strings"
)

func main() {
	isOrNot := isPalindrome("race a car")
	fmt.Print(isOrNot)
}

func isPalindrome(s string) bool {
	str := ""
	for i := 0; i < len(s); i++ {
		if isAlOrNum(s[i]) {
			str += string(s[i])
		}
	}
	lower := strings.ToLower(str)
	i, j := 0, len(lower)-1
	for i < j {
		if lower[i] == lower[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func isAlOrNum(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9'
}

func isPalindromeNum(x int) bool {
	str := fmt.Sprint(x)
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}
