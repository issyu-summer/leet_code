package main

func main() {
	twoSum2([]int{3, 2, 4}, 6)
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, num := range nums {
		m[num] = i
	}
	for i := 0; i < len(nums); i++ {
		_, ok := m[target-nums[i]]
		if ok {
			if i == m[target-nums[i]] {
				continue
			}
			return []int{i, m[target-nums[i]]}
		}
	}
	return []int{}
}
