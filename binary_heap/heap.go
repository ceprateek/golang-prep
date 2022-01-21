package binary_heap

import "fmt"

/*

10, 15, 30, 40, ,50, 100, 40


         15       30
       40  50  100  40

in := []int{1, 12, 9, 5, 6, 10, 11, 8, 7}

             12
       8            10
     5    6        9   11
   1   7

*/

type Heap []int

func (h *Heap) Heapify() {
	for i := len(*h) / 2; i >= 0; i-- {
		h.heapifyHelper(i, len(*h))
	}
}

func (h *Heap) heapifyHelper(root, len int) {
	in := *h
	left := 2*root + 1
	right := 2*root + 2
	max := root

	if left < len && in[left] > in[max] {
		max = left
	}
	if right < len && in[right] > in[max] {
		max = right
	}
	if max != root {
		in[max], in[root] = in[root], in[max]
		h.heapifyHelper(max, len)
	}
}

func (h *Heap) Del() int {
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	val := (*h)[len((*h))-1]
	*h = (*h)[:len(*h)-1]
	//balance the tree
	h.heapifyHelper(0, len(*h))
	return val
}

func (h *Heap) Insert(val int) {
	*h = append(*h, val)
	h.HeapifyAtIdx(len(*h) - 1)
}

func (h *Heap) HeapifyAtIdx(idx int){
	child := idx
	arr := *h
	parent := (child-1)/2
	if arr[child] > arr[parent]{
		arr[child], arr[parent] = arr[parent], arr[child]
		h.HeapifyAtIdx(parent)
	}
}

func PlayHeap() {
	in := Heap{1, 12, 9, 5, 6, 10, 11, 8, 7}
	in.Heapify()
	fmt.Println(in)
	val := in.Del()
	fmt.Println(val)
	fmt.Println(in)
	in.Insert(12)
	fmt.Println(in)
}
