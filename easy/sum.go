package main

import "fmt"

func main() {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var odd, even int
	for _, val := range inputList {
		if val%2 == 0 {
			even += val
		} else {
			odd += val
		}
	}
	result := []int{odd, even}
	fmt.Println(result)
}

0,1,2,3,4 := n*(n+1)/2

func missingNumber(nums []int) (result int) {
	n := len(nums)
	sum := n*(n-1)/2
	realSum := 0
	for _, val := range nums{
		realSum += val
	}
	result = sum - realSum
	return result
}
