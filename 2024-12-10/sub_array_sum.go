package _024_12_10

func subArraySum(nums []int, k int) int {
	prefixSum := func(arr []int) []int {
		//prefix[0]=0
		prefix := make([]int, len(arr)+1)
		for i := 1; i <= len(arr); i++ {
			prefix[i] = prefix[i-1] + arr[i-1]
		}
		return prefix
	}
	//l,r从1开始,所以idx为l-1,r-1
	rangeSum := func(prefix []int, l, r int) int {
		return prefix[r] - prefix[l-1]
	}
	prefix := prefixSum(nums)
	ans := 0
	for i := 1; i <= len(nums); i++ {
		for j := i; j > 0; j-- {
			if rangeSum(prefix, i, j) == k {
				ans++
			}
		}
	}
	return ans
}

func subArraySum1(nums []int, k int) int {
	prefixSum := func(arr []int) []int {
		//prefix[0]=0
		prefix := make([]int, len(arr)+1)
		for i := 1; i <= len(arr); i++ {
			prefix[i] = prefix[i-1] + arr[i-1]
		}
		return prefix
	}
	prefix := prefixSum(nums)
	ans := 0
	m := map[int]int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[prefix[i]-k]; ok {
			ans += m[prefix[i]-k]
		}
		m[prefix[i]]++
	}
	return ans
}
