package divide_conquer

import "fmt"

/*
Search Element in a sorted and rotated Array and Find the Pivot Where it is rotated
Input  : arr[] = {5, 6, 7, 8, 9, 10, 1, 2, 3};
         key = 3
Output : Found at index 8
*/

func findPivot(in []int, start, end int) int {
	if len(in) <= 1 || start < 0 || end >= len(in) || in[start] < in[end] {
		return -1
	}
	mid := start + (end-start)/2
	if in[mid] > in[mid+1] {
		return mid
	}
	p := findPivot(in, start, mid)
	if p != -1 {
		return p
	}
	p = findPivot(in, mid+1, end)
	return p
}

func PlayFindPivot() {
	in := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	fmt.Println(findPivot(in, 0, len(in)-1))
}
