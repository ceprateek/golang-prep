package greedy

import "fmt"

func maximumEvenSplit(finalSum int64) []int64 {
	var out []int64
	if finalSum%2 != 0 {
		return out
	}
	var sum, element int64
	for sum < finalSum {
		element += 2
		sum += element
		out = append(out, element)
	}
	if sum == finalSum{
		return out
	}
	diff := sum - finalSum
	idxToRemove := diff/2 - 1
	out = append(out[:idxToRemove], out[idxToRemove+1:]...)
	return out
}

func PlayMaxSplitPositiveNumbers() {
	a := maximumEvenSplit(12)
	fmt.Println(a)
}
