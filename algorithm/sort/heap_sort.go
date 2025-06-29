package main

import (
	"fmt"
)

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

func heapSort(nums []int) {
	n := len(nums)
	//为什么从 n/2 - 1 开始？	后半部分都是叶子节点，无需堆化
	//为何从下往上？确保子节点先被堆化,父节点调整一次即可
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	fmt.Println(n/2-1, nums)
	//大的放后面（）大根堆升序
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}
