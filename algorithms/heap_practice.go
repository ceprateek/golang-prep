package algorithms

import "fmt"

func PlayMyHeap() {
	in := []int{6, 4, 8, 19, 12, 11, 1}
	buildMaxHeap(&in)
	fmt.Println(in)

}

func heapifyP(nums *[]int, root, start, end int) {
	largest := root
	left := (root * 2) + 1
	right := (root * 2) + 2

	if left >= start && (*nums)[left] > (*nums)[largest] {
		largest = left
	}
	if right < end && (*nums)[right] > (*nums)[largest] {
		largest = right
	}
	if largest != root {
		(*nums)[largest], (*nums)[root] = (*nums)[root], (*nums)[largest]
		heapifyP(nums, largest, start, end)
	}
}

func buildMaxHeapP(nums *[]int) {
	for i := len(*nums); i >= 0; i-- {
		heapifyP(nums, i, 0, len(*nums))
	}
}

func insertToHeapP(nums *[]int, val int){

}

func pop(nums *int[]) int{
	top := (*nums)[0]
	
	return 0
}