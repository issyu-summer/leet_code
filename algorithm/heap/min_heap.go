package main

import (
	"fmt"
)

func main() {
	var h = &minHeap{}
	nums := []int{1, 3, 5, 2, 6, 6, 9, 10, 299, 323, 22, 345}
	//topK问题:O(nlogk),k远小于n时更优
	for i := 0; i < len(nums); i++ {
		if i < 5 {
			h.push(nums[i])
		} else if nums[i] > h.peek() {
			val, _ := h.pop()
			fmt.Println(val)
			h.push(nums[i])
		}
	}
	for i := 0; i < 5; i++ {
		val, _ := h.pop()
		fmt.Printf("%d,", val)
	}
}

type minHeap []int

func (m *minHeap) peek() int {
	return (*m)[0]
}

func (m *minHeap) push(val int) {
	*m = append(*m, val)
	m.siftUp(len(*m) - 1)
}

// push len(nums)-1,child want to up
func (m *minHeap) siftUp(i int) {
	parent := (i - 1) / 2
	if parent < 0 || (*m)[parent] <= (*m)[i] {
		return
	}
	(*m)[parent], (*m)[i] = (*m)[i], (*m)[parent]
	m.siftUp(parent)
}

func (m *minHeap) pop() (int, bool) {
	if len(*m) == 0 {
		return -1, false
	}
	(*m)[0], (*m)[len(*m)-1] = (*m)[len(*m)-1], (*m)[0]
	top := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	m.siftDown(0)
	return top, true
}

// pop,swap 0,len(nums)-1->down 0,parent want to down
func (m *minHeap) siftDown(i int) {
	l, r, smallest := 2*i+1, 2*i+2, i
	if l < len(*m) && (*m)[smallest] > (*m)[l] {
		smallest = l
	}
	if r < len(*m) && (*m)[smallest] > (*m)[r] {
		smallest = r
	}
	if smallest == i {
		return
	}
	(*m)[smallest], (*m)[i] = (*m)[i], (*m)[smallest]
	m.siftDown(smallest)
}
