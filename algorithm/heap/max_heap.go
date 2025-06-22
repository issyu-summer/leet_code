package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 2, 6, 6, 9, 10, 299, 323, 22, 345}
	h := maxHeap{[]int{}}
	for i := 0; i < len(nums); i++ {
		h.push(nums[i])
	}
	for i := 0; i < len(nums); i++ {
		val, _ := h.pop()
		fmt.Printf("%d,", val)
	}
}

type maxHeap struct {
	data []int
}

//	  1
//  /  \
// 2	3

// 0,1,2,
// 1,2,3,4,5,6,7,8
func (m *maxHeap) left(i int) int {
	return 2*i + 1
}

func (m *maxHeap) right(i int) int {
	return 2*i + 2
}

func (m *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (m *maxHeap) peek() int {
	return m.data[0]
}

func (m *maxHeap) push(val int) {
	m.data = append(m.data, val)
	m.siftUp(len(m.data) - 1)
}

func (m *maxHeap) siftUp(i int) {
	parent := m.parent(i)
	if parent < 0 || m.data[parent] >= m.data[i] {
		return
	}
	m.data[parent], m.data[i] = m.data[i], m.data[parent]
	m.siftUp(parent)
}

func (m *maxHeap) pop() (int, bool) {
	if len(m.data) == 0 {
		return -1, false
	}
	m.data[0], m.data[len(m.data)-1] = m.data[len(m.data)-1], m.data[0]
	val := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.siftDown(0)
	return val, true
}

func (m *maxHeap) siftDown(i int) {
	l, r, largest := m.left(i), m.right(i), i
	if l < len(m.data) && m.data[l] > m.data[largest] {
		largest = l
	}
	if r < len(m.data) && m.data[r] > m.data[largest] {
		largest = r
	}
	if largest == i {
		return
	}
	m.data[largest], m.data[i] = m.data[i], m.data[largest]
	m.siftDown(largest)
}
