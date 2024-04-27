package main

func main() {

}

func maxArea(height []int) int {
	i, j, m := 0, len(height)-1, 0
	for i < j {
		m = max(m, (j-i)*min(height[j], height[i]))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return m
}
