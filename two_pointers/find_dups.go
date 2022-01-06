package two_pointers

import "fmt"

//TODO

/*
Input : n = 7 and array[] = {1, 2, 3, 6, 3, 6, 1}
Output: 1, 3, 6
*/

func PlayFindDups() {
	array := []int{1, 2, 3, 6, 3, 6, 1}
	findDups(array)
}

func findDups(in []int) {
	fmt.Println(in)
	fmt.Println(len(in))
	for i := 0; i < len(in); i++ {
		in[in[i]%len(in)] +=  len(in)
	}
	fmt.Println(in)
	for i := 0; i < len(in); i++ {
		if in[i] >= len(in)* 2 {
			fmt.Println(i)
		}
	}
}
