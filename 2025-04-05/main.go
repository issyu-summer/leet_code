package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	res := sort.Search(len(nums), func(i int) bool {
		fmt.Println("idx", i, "val", nums[i])
		if (nums[i] >= nums[0]) == (target >= nums[0]) {
			return nums[i] >= target
		}
		//nums[i]和target在不同段
		return !(nums[i] >= nums[0])
	})
	if res < len(nums) && nums[res] == target {
		return res
	}
	return -1
}

func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	//f(i)=true向左，f(i)=false向右
	res := sort.Search(len(nums), func(i int) bool {
		return nums[i] < nums[0]
	})
	if res < len(nums) {
		return nums[res]
	}
	return nums[0]
}

func isValid(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) != 0 && m[stack[len(stack)-1]] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.min[len(this.stack)-1]
	this.min = append(this.min, min(val, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
