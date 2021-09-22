package arrays

import "fmt"

/*
{34, 8, 10, 3, 2, 80, 30, 33, 1}

*/

func PlayDiffMax() {
	in := []int{34, 8, 10, 3, 2, 80, 30, 33, 1}
	fmt.Println(findMaxIndexDifference(in))
}

func findMaxIndexDifference(in []int) (r, l, diff int) {
	for i := 0; i < len(in)-1; i++ {
		for j := len(in) - 1; j > i; j-- {
			if in[j] > in[i] && j-i > diff {
				diff = j - i
				r = j
				l = i
			}
		}
	}
	return r, l, diff
}

func findMaxDifferenceEfficient(int []int) (r, l, diff int) {

	return r, l, diff
}
