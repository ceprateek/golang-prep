package arrays

import "fmt"

func PlayMaxCircularSum() {
	input := []int{11, 1, -17, 2, -15, 9, 13}
	max := findMaxCircularSum(input)
	fmt.Println(max)
}

func findMaxCircularSum(input []int) (maxSum int) {
	maxSumSoFar := 0
	i := 0
	lastIndex := len(input)
	for i < lastIndex {
		maxSumSoFar = maxSumSoFar + input[i]
		if maxSumSoFar < 0 {
			maxSumSoFar = 0
		}
		if maxSumSoFar > maxSum {
			maxSum = maxSumSoFar
		}
		i++
		if i == len(input) {
			i = 0
			lastIndex = len(input) - 1
		}
	}
	return maxSum
}
