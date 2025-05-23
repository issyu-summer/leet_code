package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next

		//从后向前进行连接
		first.Next = second.Next
		second.Next = first
		pre.Next = second

		pre = first
	}
	return dummy.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	var decode func() *TreeNode
	decode = func() *TreeNode {
		s := nodes[0]
		nodes = nodes[1:]
		if s == "nil" {
			return nil
		}
		val, _ := strconv.Atoi(s)
		return &TreeNode{
			Val:   val,
			Left:  decode(),
			Right: decode(),
		}
	}
	return decode()
}

func bfsSerialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	q := []*TreeNode{root}
	var res string
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if node == nil {
				res += "nil,"
				continue
			}
			res += fmt.Sprintf("%d,", node.Val)
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}
	return res
}

func bfsDeserialize(data string) *TreeNode {
	if data == "nil" {
		return nil
	}
	nodes := strings.Split(data, ",")
	for nodes[len(nodes)-1] == "nil" || nodes[len(nodes)-1] == "" {
		nodes = nodes[:len(nodes)-1]
	}
	s := nodes[0]
	val, _ := strconv.Atoi(s)
	root := &TreeNode{Val: val}
	q := []*TreeNode{root}
	var index = 1
	for len(q) > 0 && index < len(nodes) {
		node := q[0]
		q = q[1:]
		if index < len(nodes) && nodes[index] != "nil" {
			val, _ := strconv.Atoi(nodes[index])
			node.Left = &TreeNode{Val: val}
			q = append(q, node.Left)
		}
		index++
		if index < len(nodes) && nodes[index] != "nil" {
			val, _ := strconv.Atoi(nodes[index])
			node.Right = &TreeNode{Val: val}
			q = append(q, node.Right)
		}
		index++
	}
	return root
}

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
