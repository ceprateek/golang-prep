package algorithms

import "fmt"

var counter int

func binarySearch(input []int, target int, start int, end int) int {
	counter++
	if start < 0 || end >= len(input) || counter > 10 {
		return -1
	}
	mid := start + ((end - start) / 2)
	fmt.Println(fmt.Sprintf("start: %d end:%d counter:%d mid:%d", start, end, counter, mid))
	if target == input[mid] {
		return mid
	} else if target < input[mid] {
		return binarySearch(input, target, start, mid-1)
	} else {
		return binarySearch(input, target, mid+1, end)
	}
}

func search(input []int, target int) int {
	return binarySearch(input, target, 0, len(input)-1)
}

func PlayBinarySearch() {
	in := []int{1, 5, 6, 7, 8, 9, 10, 11, 12}
	target := 12
	fmt.Println(search(in, target))
}
