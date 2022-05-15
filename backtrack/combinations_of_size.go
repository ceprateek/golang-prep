package backtrack

import "fmt"

func PlayFindCombinations(){
	r := findCombinations(5,3)
	fmt.Println(r)
}

func findCombinations(n, k int) [][]int {
	result := make([][]int, 0, (n-1)*k)
	temp := make([]int, 0, k)
	findCombinationsHelper(n, k, &temp, 1, &result)
	return result
}

func findCombinationsHelper(n, k int, temp *[]int, start int, result *[][]int) {
	if len(*temp) == k {
		a := make([]int, k)
		copy(a, *temp)
		*result = append(*result, a)
	} else {
		for i:=start;i<=n;i++{
			*temp = append(*temp, i)
			findCombinationsHelper(n,k,temp,i+1, result)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}
