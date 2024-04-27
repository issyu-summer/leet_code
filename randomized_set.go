package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	nums []int
	idx  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		idx:  map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.idx[val]
	if ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.idx[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.idx[val]
	if !ok {
		return false
	}
	id := this.idx[val]
	las := len(this.nums) - 1
	this.nums[id] = this.nums[las]
	this.idx[this.nums[id]] = id
	this.nums = this.nums[:las]
	delete(this.idx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
