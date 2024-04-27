package main

import "fmt"

func main() {
	result := isValid("()[]{}")
	fmt.Println(result)
}

var dict = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	stack := []byte{'?'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == dict[stack[len(stack)-1]] ||
			s[i] == dict[stack[len(stack)-1]] ||
			s[i] == dict[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 1
}
