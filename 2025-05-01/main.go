package main

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	//单调队列
	var q, res []int
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && q[0] < i-k+1 {
			q = q[1:]
		}
		//1,3,-1,-3,5,3,6,7
		//单调递减队列，因此idx是q[0]的是最小值
		for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}
