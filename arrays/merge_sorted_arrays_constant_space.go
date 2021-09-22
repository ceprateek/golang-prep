package arrays

import "fmt"

/*
input : [3 4 5 7 9] [1 6 8]
output: [1 3 4 5 6]  [7 8 9]
*/

/*
[1 4 5 7 9] [2 6 8]

*/

func PlayMergerArrays(){
	in := []int{3,4,5,9, 11}
	in1 := []int{1,6,8}
	mergeTwoArrays(&in, &in1)
	fmt.Println(in)
	fmt.Println(in1)
}

func mergeTwoArrays(in1, in2 *[]int) {
	A := *in1
	B := *in2
	for i := 0; i < len(A); i++ {
		if A[i] > B[0] {
			A[i], B[0] = B[0], A[i]
			for j := 0; j < len(B)-1; j++ {
				if B[j] > B[j+1] {
					B[j], B[j+1] = B[j+1], B[j]
				}
			}
		}
	}
}
