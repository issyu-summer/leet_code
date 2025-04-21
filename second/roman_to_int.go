package main

func main() {

}

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	num := romanMap[s[0]]
	for _, b := range s[1:] {
		nextNum := romanMap[byte(b)]
		if num < nextNum {
			sum -= num
		} else {
			sum += num
		}
		num = nextNum
	}
	return sum + num
}
