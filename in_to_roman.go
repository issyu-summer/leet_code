package main

func main() {

}

var m = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) string {
	sortedValue := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for _, val := range sortedValue {
		for num >= val {
			num -= val
			result += m[val]
		}
		if num == 0 {
			break
		}
	}
	return result
}
