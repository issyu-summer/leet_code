package main

import (
	"bytes"
	"math"
)

func main() {

}

// 最长有效括号
func longestValidParentheses(s string) int {
	var stack []int
	var res int
	for r := 0; r < len(s); r++ {
		if s[r] == '(' {
			stack = append(stack, r)
		} else {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				var curLen int
				if len(stack) == 0 {
					curLen = r - 0 + 1
				} else {
					curLen = r - (stack[len(stack)-1] + 1) + 1
				}
				res = max(res, curLen)
			} else {
				stack = append(stack, r)
			}
		}
	}
	return res
}

// 有效的括号
func isValid(s string) bool {
	var stack []int
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 && m[s[stack[len(stack)-1]]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// MinStack 最小栈
type MinStack struct {
	stack []int
	minus []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minus: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.minus[len(this.minus)-1]
	this.minus = append(this.minus, min(top, val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minus = this.minus[:len(this.minus)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minus[len(this.minus)-1]
}

// 字符串解码
func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			var str []byte
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				str = append([]byte{stack[len(stack)-1]}, str...)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			//逆向数字
			num, base := 0, 1
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				num += int(stack[len(stack)-1]-'0') * base
				stack = stack[:len(stack)-1]
				base *= 10
			}
			stack = append(stack, bytes.Repeat(str, num)...)
		}
	}
	return string(stack)
}

// 每日温度，单调递减栈
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		//73,74,75,71,69,72,76
		//73
		//74
		//75,71,69
		//75,72
		//76
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIdx := stack[len(stack)-1]
			res[prevIdx] = i - prevIdx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

// 柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	//2,1,5,6,2,3
	//2(2*1=2)
	//1,5,6(6*1=6)
	//1,5(5*2=10)
	//1(1*3=3)
	var res int
	var stack []int
	heights = append(heights, 0)
	for r := 0; r < len(heights); r++ {
		for len(stack) > 0 && heights[r] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[top]
			var width int
			if len(stack) == 0 {
				width = r - (-1 + 1)
			} else {
				width = r - (stack[len(stack)-1] + 1)
			}
			res = max(res, width*height)
		}
		stack = append(stack, r)
	}
	return res
}
