package algorithms

import (
	"fmt"
	"math"
)

//Input:  { 4, -1, 2, 8, -5, 4}
//Output: Subarray with the largest sum is {4, -1, 2, 1} with sum 6.

func FindMaxSumContinuous(input []int) (maxSum int, start int, end int) {
	maxSum = math.MinInt32
	currentMaxSum := 0
	s := 0
	for i, val := range input {
		currentMaxSum = currentMaxSum + val
		if currentMaxSum > maxSum {
			maxSum = currentMaxSum
			end = i
			start = s
		}
		if currentMaxSum < 0 {
			currentMaxSum = 0
			s = i + 1
		}
	}
	return maxSum, start, end
}

func PlayMaxSumSubarray() {
	in := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	sum, start, end := FindMaxSumContinuous(in)
	fmt.Println(fmt.Sprintf("start: %d end: %d sum: %d", start, end, sum))
}
