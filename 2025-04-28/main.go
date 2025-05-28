package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(maxGain(root.Left), 0)
		right := max(maxGain(root.Right), 0)
		res = max(left+right+root.Val, res)
		//只能返回一支，否则会分叉
		return root.Val + max(left, right)
	}
	maxGain(root)
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	hasCycle := false
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	res := head
	for res != slow {
		res = res.Next
		slow = slow.Next
	}
	return slow
}

func restoreIpAddresses(s string) []string {
	var res []string
	if len(s) < 4 || len(s) > 12 {
		return res
	}
	valid := func(s string) bool {
		//先导0
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		num, _ := strconv.Atoi(s)
		valid := num >= 0 && num <= 255
		fmt.Println("str", s, "valid", valid)
		return valid
	}
	var backTrack func(path []string, remain string, k int)
	backTrack = func(path []string, remain string, k int) {
		fmt.Println("remain", remain, "k", k)
		if k == 4 && remain == "" {
			res = append(res, strings.Join(path, "."))
			return
		}
		//每个segment最少取1个字符，最多取3个字符
		for i := 1; i <= 3; i++ {
			//剩余的字符不够选取了,则跳过
			//选取的字符不合规则, 则跳过
			fmt.Println(i, remain)
			if i > len(remain) {
				continue
			}
			selectStr := remain[:i]
			if !valid(selectStr) {
				continue
			}
			//k是选取的次数
			if k < 3 {
				remainLen := len(remain) - i
				//第一次选取，k=0，则如果剩余的长度<1*(3-0)，小3个字符，则说明不足以选择剩余的1，2，3次，所以跳过
				//同理，如果剩余的字符>3*(3-0)，大于9个字符，则说明剩余的1，2，3每次选择3个，还是有剩余，则跳过
				if remainLen < (3-k) || remainLen > 3*(3-k) {
					continue
				}
				//k=3时，是最后一次选择，则必须全用完才行，否则会有剩余字符
			} else {
				if i != len(remain) {
					continue
				}
			}
			backTrack(append(path, selectStr), remain[i:], k+1)
		}
	}
	backTrack([]string{}, s, 0)
	return res
}

func longestCommonSubsequence(text1 string, text2 string) int {
	f := make([][]int, len(text1)+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[len(text1)][len(text2)]
}
