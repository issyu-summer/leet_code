package main

import (
	"container/list"
	"strings"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < cnt; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func subsets(nums []int) [][]int {
	var res [][]int
	var backTracK func(path []int, start int)
	backTracK = func(path []int, start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			backTracK(append(path, nums[i]), i+1)
		}
	}
	backTracK([]int{}, 0)
	return res
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		elem.Value.(*entry).val = value
		this.list.MoveToFront(elem)
		return
	}
	if this.list.Len() >= this.capacity {
		back := this.list.Back()
		delete(this.cache, back.Value.(*entry).key)
		this.list.Remove(back)
	}
	elem := this.list.PushFront(&entry{key, value})
	this.cache[key] = elem
}

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	reverse := func(str []string) {
		i, j := 0, len(str)-1
		for i < j {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
	reverse(arr)
	var filter []string
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			filter = append(filter, arr[i])
		}
	}
	return strings.Join(filter, " ")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	//preorder是根、左、右
	i := 0
	for i < len(inorder) {
		if inorder[i] == preorder[0] {
			break
		}
		i++
	}
	// i是inorder中的根节点，左、根、右，所以，左子树是[:i],右子树是[i+1:]
	leftInOrder := inorder[:i]
	rightInOrder := inorder[i+1:]
	// preorder的左子树、右子树与其长度有关系，
	// 左子树：0(根近点)->1(根节点)->推倒右子树r-l+1=len=>len+l-1=>len,但是左闭右开=>len+1
	leftPreOrder := preorder[1 : len(leftInOrder)+1]
	rightPreOrder := preorder[len(leftInOrder)+1:]
	root := &TreeNode{preorder[0], nil, nil}
	root.Left = buildTree(leftPreOrder, leftInOrder)
	root.Right = buildTree(rightPreOrder, rightInOrder)
	return root
}
