package bit

import (
	"fmt"
)

/*
You are given a list of numbers where every number occurs twice except one, you have to find the number that appeared only once.
Input: [ 5 3 9 11 5 11 9 ]
Output: 3
*/

func PlayFindSingleNumber() {
	fmt.Println(findSingleNumber([]int{5, 3, 9, 11, 5, 11, 9}))
}

func findSingleNumber(in []int) (out int) {
	for _, val := range in {
		out = out ^ val
		fmt.Println(out)
	}
	return out
}
