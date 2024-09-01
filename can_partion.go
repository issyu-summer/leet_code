package main

func main() {

}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	f := make([]bool, target+1)
	f[0] = true
	//外层循环物品
	for _, num := range nums {
		if num > target {
			return false
		}
		if f[target-num] {
			return true
		}
		//内层循环容量
		for i := target; i >= num; i-- {
			f[i] = f[i] || f[i-num]
		}
	}
	return f[target]
}
