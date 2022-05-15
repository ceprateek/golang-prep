package recursion

import "fmt"

/*
https://leetcode.com/problems/permutations/
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

sol:
[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]

prepend 1 with permutations of 2,3
prepend 2 with permutations of 1,3
prepend 3 with permutations of 1,1
*/

func PlayPermute(){
	in := []int{1,2,3}
	r := permute(in)
	fmt.Println(r)
}

func permute(in []int) [][]int {
	var result [][]int
	t := make([]int, 0, len(in))
	permuteHelper(&in, &result, &t)
	return result
}

func permuteHelper(in *[]int, result *[][]int, temp *[]int) {
	if len(*in) == 0 {
		a := make([]int, len(*temp))
		copy(a, *temp)
		*result = append(*result, a)
		t := make([]int, 0, len(*in))
		temp = &t
		return
	}
	for i := 0; i < len(*in); i++ {
		chosen := (*in)[i]
		*temp = append(*temp, chosen)
		*in =append((*in)[:i], (*in)[i+1:]...)
		permuteHelper(in, result, temp)
		*temp = (*temp)[:len(*temp)-1]
		*in = append((*in)[:i+1], (*in)[i:]...)
		(*in)[i] = chosen
	}
}
