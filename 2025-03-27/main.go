package main

func main() {
	reverseBetween(parseListToLink([]int{1, 2, 3, 4, 5}), 2, 4)
}

func parseListToLink(ar []int) *ListNode {
	pre := &ListNode{}
	cur := pre
	for _, i := range ar {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, l, r int) *ListNode {
	pre := &ListNode{Next: head}
	lHP, rTP := pre, pre
	for i := 0; i < l-1; i++ {
		lHP = lHP.Next
	}
	for i := 0; i < r; i++ {
		rTP = rTP.Next
	}
	lh := lHP.Next
	lHP.Next = nil
	rt := rTP.Next
	rTP.Next = nil
	reverse(lh)
	lHP.Next = rTP
	lh.Next = rt
	return pre.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func findKthLargest(nums []int, k int) int {

}

func partition(nums []int, low, high int) int {
	//base num
	pivot := nums[low]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func quickSortInPlace(nums []int, low, high int) {
	if low < high {
		pi := partition(nums, low, high)
		quickSortInPlace(nums, low, pi-1)
		quickSortInPlace(nums, pi+1, high)

	}
}

func quickSort(nums []int) []int {
	quickSortInPlace(nums, 0, len(nums)-1)
	return nums
}
