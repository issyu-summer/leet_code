package _024_12_02

func towSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		_, ok := m[num]
		if ok {
			return []int{
				m[num], i,
			}
		}
		m[target-num] = i
	}
	return []int{}
}
