package main

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j, tail := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[tail] = nums1[i]
			i--
			tail--
		} else {
			nums1[tail] = nums2[j]
			j--
			tail--
		}
	}
	for i >= 0 {
		nums1[tail] = nums1[i]
		tail--
		i--
	}
	for j >= 0 {
		nums1[tail] = nums2[j]
		tail--
		j--
	}
}
