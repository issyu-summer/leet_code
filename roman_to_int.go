package main

func main() {

}

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var (
		sum   = 0
		bytes = []byte(s)
		num   = romanMap[bytes[0]]
	)
	for i := 1; i < len(bytes); i++ {
		nextNum := romanMap[bytes[i]]
		if num < nextNum {
			sum -= num
		} else {
			sum += num
		}
		num = nextNum
	}
	sum += num
	return sum
}
