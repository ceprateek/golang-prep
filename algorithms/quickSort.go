package algorithms

import "fmt"

func Quicksort() {
	arr := &[]int{12, 3, 18, 24, 0, 5, 2}
	quickSortHelper(arr, 0, len(*arr)-1)
	fmt.Println(*arr)
}

func quickSortHelper(in *[]int, start int, end int) {
	if start>=end{
		return
	}
	pivot := partition(in, start, end)
	quickSortHelper(in, start, pivot-1)
	quickSortHelper(in, pivot+1, end)

}

//end should be length-1 or less
func partition(in *[]int, start int, end int) int {
	pivot := (*in)[end]
	pindex := start
	for i := start; i < end; i++ {
		if (*in)[i] < pivot {
			(*in)[i], (*in)[pindex] = (*in)[pindex], (*in)[i]
			pindex++
		}
	}
	(*in)[end], (*in)[pindex] = (*in)[pindex], (*in)[end]
	return pindex
}
