package main

func main() {
}

// 归并排序并非原地排序
func mergeSort(nums []int) []int {
	//r-l+1(len)<=1=>l>=r
	if len(nums) <= 1 {
		return nums
	}
	//r= len(nums) l = 0
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
