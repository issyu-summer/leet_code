package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

func main() {
	//5,5,5
	//5,5,2
	//5,5,1
	//5,2,5
	//
	fmt.Println(multiply("2", "3"))
}

func coinChange(coins []int, amount int) int {
	var f func(coins []int, amount int) int
	var memo = map[int]int{}
	f = func(coins []int, amount int) int {
		if amount == 0 {
			return 0
		}
		if cnt, ok := memo[amount]; ok {
			return cnt
		}
		res := math.MaxInt32
		for i := 0; i < len(coins); i++ {
			if coins[i] > amount {
				continue
			}
			cnt := f(coins, amount-coins[i])
			res = min(res, cnt+1)
		}
		if res == math.MaxInt32 {
			res = -1
		}
		memo[amount] = res
		return res
	}
	return f(coins, amount)
}

func coinChangeUseDp(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 0; i < len(f); i++ {
		f[i] = math.MaxInt32
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] = min(f[j], f[j-coins[i]]+1)
		}
	}
	if f[amount] == math.MaxInt32 {
		return -1
	}
	return f[amount]
}

func firstMissingPositive(nums []int) int {
	var i int
	for i < len(nums) {
		//原地hash
		if nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			return j + 1
		}
	}
	return len(nums) + 1
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	ans := ""
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		ans += strconv.Itoa(res[i])
	}
	return ans
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCache) Put(key int, value int) {
	if elem, ok := l.cache[key]; ok {
		elem.Value.(*entry).val = value
		l.list.MoveToFront(elem)
		return
	}
	if l.list.Len() >= l.capacity {
		elem := l.list.Back()
		delete(l.cache, elem.Value.(*entry).key)
		l.list.Remove(elem)
	}
	elem := l.list.PushFront(&entry{key, value})
	l.cache[key] = elem
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(entry).val
	}
	return -1
}
