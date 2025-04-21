package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	startIndex, currGas := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currGas += gas[i] - cost[i]
		if currGas < 0 {
			startIndex = i + 1
			currGas = 0
		}
	}
	if totalGas < totalCost {
		return -1
	}
	return startIndex
}
