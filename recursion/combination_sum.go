package recursion

import "fmt"

/*
https://leetcode.com/problems/combination-sum/
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
*/

func PlayCombinationSum(){
	in := []int{2,3,6,7}
	result := make([][]int, 0)
	sum := 7
	var tempResult []int
	var tempsum int
	findCombinationSumHelper(&in, sum, &result, &tempResult, &tempsum)
	fmt.Println(result)
}

func indent(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("   ")
	}
}

func findCombinationSumHelper(in *[]int, sum int, result *[][]int,
	tempResult *[]int, tempSum *int) int{
	indent(len(*tempResult))
	fmt.Println(fmt.Sprintf("input: %v tempResult: %v", *in, *tempResult))
	if *tempSum == sum {
		a := make([]int, len(*tempResult))
		copy(a, *tempResult)
		*result = append(*result, a)
		return *tempSum
	} else {
		for i := 0; i < len(*in); i++ {
			//choose
			choose := (*in)[i]
			*in =append((*in)[:i], (*in)[i+1:]...)
			tempResult1 := append(*tempResult, choose)
			newSum := *tempSum + choose
			for newSum <=sum {
				indent(len(*tempResult))
				fmt.Println(fmt.Sprintf("choose: %v", choose))
				findCombinationSumHelper(in, sum, result, &tempResult1, &newSum)
				newSum += choose
				tempResult1 = append(tempResult1, choose)
			}

			//un-choose
			*in = append((*in)[:i+1], (*in)[i:]...)
			(*in)[i] = choose
			indent(len(*tempResult))
			fmt.Println(fmt.Sprintf("unchoose c: %v input: %v temp: %v", choose, *in, *tempResult))
		}
	}
	return *tempSum
}
