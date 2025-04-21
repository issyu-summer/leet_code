package main

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, n int, nums2 []int, m int) {
	if n == 0 {
		copy(nums1, nums2)
	}
	if m == 0 {
		return
	}
	i, j := 0, 0
	ans := make([]int, 0, m+n)
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	if i != n {
		ans = append(ans, nums1[i:]...)
	}
	if j != m {
		ans = append(ans, nums2[j:]...)
	}
	copy(nums1, ans)
}

func mergeII(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
	for t := 0; t <= j; t++ {
		nums1[t] = nums2[t]
	}
}

func handleRemain() {
	// handle remaining part
}

func IsomorphicPointerTemplate(nums1, nums2 []int) {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums1) {
		do()
	}
	handleRemain()
}
