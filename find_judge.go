package main

func main() {

}

func findJudge(n int, trust [][]int) int {
	//1~n
	inDegrees, outDegrees := make([]int, n+1), make([]int, n+1)
	for _, t := range trust {
		inDegrees[t[1]]++
		outDegrees[t[0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegrees[i] == n-1 && outDegrees[i] == 0 {
			return i
		}
	}
	return -1
}
