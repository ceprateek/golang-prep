package greedy

import "fmt"

func PlayCanComplete() {
	gas := []int{3,1,1}
	cost := []int{1,2,2}
	result := canCompleteCircuit(gas, cost)
	fmt.Println(result)
}

func canCompleteCircuit(gas []int, cost []int) int {
	var totalGas, startPosition, currentTank int
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		currentGas := gas[i] - cost[i]
		currentTank += currentGas
		if currentTank < 0 {
			startPosition = i + 1
			currentTank = 0
		}
	}
	if totalGas >= 0 && startPosition < len(gas) {
		return startPosition
	} else {
		return -1
	}
}

/*

3   1   1
1   2   2
2   -1  -1
*/
