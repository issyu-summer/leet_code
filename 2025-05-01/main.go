package main

import (
	"container/list"
	"sort"
	"strconv"
	"strings"
)

func main() {

}

func generateParenthesis(n int) []string {
	var res []string
	var backTrack func(path string, left, right int)
	backTrack = func(path string, left, right int) {
		if len(path) == n*2 {
			res = append(res, path)
			return
		}
		if left < n {
			backTrack(path+"(", left+1, right)
		}
		if right < left {
			backTrack(path+")", left, right+1)
		}
	}
	backTrack("", 0, 0)
	return res
}

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	r := n - 2
	for r >= 0 && nums[r] >= nums[r+1] {
		r--
	}
	//非全部降序
	if r >= 0 {
		i := n - 1
		for nums[i] <= nums[r] {
			i--
		}
		nums[i], nums[r] = nums[r], nums[i]
	}
	reverse(nums, r+1, len(nums)-1)
}
func reverse(nums []int, start, end int) {
	l, r := start, end
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func mySqrt(x int) int {
	res := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if res*res == x {
		return res
	} else if res*res > x {
		return res - 1
	}
	return -1
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		var num1, num2 int
		if i < len(v1) {
			num1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			num2, _ = strconv.Atoi(v2[i])
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

// LRUCache LRU Cache练习
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (l *LRUCache) Put(key, val int) {
	//存在则更新
	if elem, ok := l.cache[key]; ok {
		elem.Value.(*entry).value = val
		l.list.MoveToFront(elem)
		return
	}
	//长度超限则淘汰back
	if l.list.Len() >= l.capacity {
		elem := l.list.Back()
		if elem != nil {
			l.list.Remove(elem)
			delete(l.cache, elem.Value.(*entry).key)
		}
	}
	//最后set
	elem := l.list.PushFront(&entry{key, val})
	l.cache[key] = elem
}
