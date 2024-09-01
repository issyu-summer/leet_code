package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	sMap := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for i := 0; i < len(str); i++ {
			cnt[str[i]-'a']++
		}
		sMap[cnt] = append(sMap[cnt], str)
	}
	var ans [][]string
	for _, ar := range sMap {
		ans = append(ans, ar)
	}
	return ans
}
