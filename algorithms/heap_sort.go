package algorithms

import (
	"fmt"
)

func HeapSort(in *[]int) {
	buildMaxHeap(in)
	len := len(*in)
	for i := len - 1; i >= 0; i-- {
		(*in)[i], (*in)[0] = (*in)[0], (*in)[i]
		heapify(in, 0, i)
	}
	fmt.Println(in)
}

func HeapSortReverse(in *[]int) {
	buildMinHeap(in)
	len := len(*in)
	for i := len - 1; i >= 0; i-- {
		(*in)[i], (*in)[0] = (*in)[0], (*in)[i]
		heapifyMin(in, 0, i)
	}
	fmt.Println(in)
}

func buildMaxHeap(in *[]int) {
	for i := len(*in)/2 - 1; i >= 0; i-- {
		heapify(in, i, len(*in))
	}
}

func heapify(in *[]int, root int, len int) {
	largest := root
	left := root*2 + 1
	right := root*2 + 2
	if left < len && (*in)[left] > (*in)[largest] {
		largest = left
	}
	if right < len && (*in)[right] > (*in)[largest] {
		largest = right
	}
	if largest != root {
		(*in)[largest], (*in)[root] = (*in)[root], (*in)[largest]
		heapify(in, largest, len)
	}
}

func buildMinHeap(in *[]int){
	for i := len(*in)/2 - 1; i >= 0; i-- {
		heapifyMin(in, i, len(*in))
	}
}

func heapifyMin(in *[]int, root int, len int) {
	lowest := root
	left := root*2 + 1
	right := root*2 + 2
	if left < len && (*in)[left] < (*in)[lowest] {
		lowest = left
	}
	if right < len && (*in)[right] < (*in)[lowest] {
		lowest = right
	}
	if lowest != root {
		(*in)[lowest], (*in)[root] = (*in)[root], (*in)[lowest]
		heapifyMin(in, lowest, len)
	}
}

func PlayHeapSort() {
	in := []int{1, 12, 9, 5, 6, 10, 11, 8, 7}
	HeapSort(&in)
	HeapSortReverse(&in)
}
