package main

func main() {

}

func canConstruct(ransomNot string, magazine string) bool {
	mMap := make(map[byte]int)
	for _, b := range []byte(magazine) {
		mMap[b]++
	}
	for _, r := range []byte(ransomNot) {
		mMap[r]--
		if mMap[r] < 0 {
			return false
		}
	}
	return true
}

func canConstructArrVersion(ransomNot string, magazine string) bool {
	if len(magazine) < len(ransomNot) {
		return false
	}
	var m [26]int
	for _, r := range magazine {
		m[r-'a']++
	}
	for _, r := range ransomNot {
		m[r-'a']--
		if m[r-'a'] < 0 {
			return false
		}
	}
	return true
}
