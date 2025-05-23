package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{1, 5, 4, 3, 6, 8, 9}
	heapSort(nums)
	fmt.Println(nums)
}

func heapSort(nums []int) {
	n := len(nums)
	//为什么从 n/2 - 1 开始？	后半部分都是叶子节点，无需堆化
	//为何从下往上？确保子节点先被堆化,父节点调整一次即可
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	fmt.Println(n/2-1, nums)
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}

// 改变堆的符号，可以变为逆序
func heapify(nums []int, heapSize, i int) {
	fmt.Printf("heapify: %v (i=%d)\n", nums, i)
	//二叉树
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	//fmt.Printf("heapify: %v (largest:%d,l:%d,r:%d)\n", nums, largest, l, r)
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	//递归构造大根堆
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, heapSize, largest)
	}
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
