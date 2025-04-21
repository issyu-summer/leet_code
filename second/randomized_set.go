package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	nums []int
	//num->idx
	idx map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		idx:  map[int]int{},
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	_, ok := r.idx[val]
	if ok {
		return false
	}
	r.nums = append(r.nums, val)
	r.idx[val] = len(r.nums) - 1
	return true
}
func (r *RandomizedSet) Remove(val int) bool {
	id, ok := r.idx[val]
	if !ok {
		return false
	}
	r.nums[id] = r.nums[len(r.nums)-1]
	r.idx[r.nums[id]] = id
	r.nums = r.nums[:len(r.nums)-1]
	//需要删除不然有两个对应位置i
	delete(r.idx, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}
