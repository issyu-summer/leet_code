package main

func main() {
	longestCommonSubsequence("abcde", "ace")

}

// 通过扩展行列避免初始状态问题
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1)+1, len(text2)+1
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	//1.init status
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				ans[i+1][j+1] = ans[i][j] + 1
			} else {
				ans[i+1][j+1] = max(ans[i+1][j], ans[i][j+1])
			}
		}
	}
	return ans[m-1][n-1]
}
