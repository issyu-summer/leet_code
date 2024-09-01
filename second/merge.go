package main

import "context"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var ans []int
	i, j := 0, 0
	for {
		if i == m {
			ans = append(ans, nums2[j:]...)
			break
		}
		if j == n {
			ans = append(ans, nums1[i:]...)
			break
		}
		if nums1[i] <= nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	copy(nums1, ans)
	background := context.Background()

}
