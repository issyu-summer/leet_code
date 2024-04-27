package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; {
		sumOfGas, sumOfCost, nextPositionWhichCantReach := 0, 0, 0
		for nextPositionWhichCantReach < n {
			j := (i + nextPositionWhichCantReach) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfGas < sumOfCost {
				break
			}
			nextPositionWhichCantReach++
		}
		if nextPositionWhichCantReach == n {
			return i
		} else {
			i += nextPositionWhichCantReach + 1
		}
	}
	return -1
}
