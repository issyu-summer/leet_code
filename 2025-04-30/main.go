package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(mergeSort([]int{math.MaxInt, math.MinInt, 9, 2, 333, 81281, 333}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var fc func(root *TreeNode)
	fc = func(root *TreeNode) {
		if root == nil {
			return
		}
		fc(root.Left)
		res = append(res, root.Val)
		fc(root.Right)
	}
	fc(root)
	return res
}

func search(nums []int, target int) int {
	idx := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if idx < len(nums) && nums[idx] == target {
		return idx
	}
	return -1
}

type MyQueue struct {
	//stack先进后出->队列先进先出=>所以需要反向
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) moveToOutStack() {
	for len(this.inStack) > 0 {
		val := this.inStack[len(this.inStack)-1]
		this.inStack = this.inStack[:len(this.inStack)-1]
		this.outStack = append(this.outStack, val)
	}
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.moveToOutStack()
	}
	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.moveToOutStack()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// dummy->1->2->3->4
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil
	fmt.Println(head, mid)
	l := sortList(head)
	r := sortList(mid)
	return merge(l, r)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return mergeArr(left, right)
}

func mergeArr(l, r []int) []int {
	var res []int
	var i, j int
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	if i < len(l) {
		res = append(res, l[i:]...)
	}
	if j < len(r) {
		res = append(res, r[j:]...)
	}
	return res
}
