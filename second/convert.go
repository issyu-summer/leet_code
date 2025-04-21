package main

func main() {

}

// numRows为3,则横轴为0，1，2(i++),1,0(i--),1,2(LOOP)
// 因此当i==0或者i==numRows-1的时候，需要翻转加数
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	ar := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		ar[i] = make([]byte, 0)
	}
	i, flag := 0, -1
	for j := 0; j < len(s); j++ {
		ar[i] = append(ar[i], s[j])
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	var ans string
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar[i]); j++ {
			ans += string(ar[i][j])
		}
	}
	return ans
}
