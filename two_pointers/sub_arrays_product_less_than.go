package two_pointers

import "fmt"

/*
Given an array of positive numbers, calculate the number of possible
contiguous subarray having product lesser than a given number K.
Input : arr[] = [1, 2, 3, 4]
        K = 10
Output : 7
*/

var counter int

func subArrayProductHelper(in, temp *[]int, k, currentK int) {
	if len(*in) == 0 {
		if currentK < k{
			counter++
		}
	} else {
		//case with the element
		//choose
		chosen := (*in)[0]
		*in = (*in)[1:]
		//explore
		subArrayProductHelper(in, temp, k, currentK)
		//un-choose
		*temp = append(*temp, chosen)
		currentK *= chosen
		//case without the element
		//explore
		subArrayProductHelper(in, temp, k, currentK)
		//un-choose
		*temp = (*temp)[:len(*temp)-1]
		currentK = currentK/chosen
		*in = insert(*in, 0, chosen)
	}
}

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func subArraysProductLessThanK(in []int, k int) {
	temp := make([]int, 0, len(in))
	initProduct := 1
	subArrayProductHelper(&in, &temp, k, initProduct)
}

func PlaySubArraysLessThanK() {
	in := []int{1,2}
	subArraysProductLessThanK(in, 10)
	fmt.Println(counter)
}
