package heap_practice

import "fmt"

type myHeap struct {
	data []int
}

func (h *myHeap) Heapify() {
	startIndex := len(h.data) - 1
	for i := startIndex; i >= 0; i-- {
		h.heapifyIndex(i)
	}
}

func (h *myHeap) heapifyIndex(i int) {
	n := len(h.data)
	largest := i
	right := i*2 + 2
	left := i*2 + 1

	if left < n && h.data[left] > h.data[largest] {
		largest = left
	}
	if right < n && h.data[right] > h.data[largest] {
		largest = right
	}
	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.heapifyIndex(largest)
	}
}

func (h *myHeap) heapifyInsert(index int){
	if index < 0{
		return
	}
	parent := (index-1)/2
	if h.data[parent] < h.data[index]{
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		h.heapifyInsert(parent)
	}
}

func (h *myHeap) Insert(val int){
	h.data = append(h.data, val)
	h.heapifyInsert(len(h.data)-1)
}

func (h *myHeap) PrintHeap() {
	fmt.Println("\nArray representation of heap is:")

	for i := 0; i < len(h.data); i++ {
		fmt.Print(h.data[i], " ")
	}
	fmt.Println()
}

func (h *myHeap) Delete() int{
	result := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyIndex(0)
	return result
}

func PlayHeap() {
	arr := []int{3, 5, 9, 6, 8, 20, 10, 12, 18, 9}
	h := myHeap{
		data: arr,
	}
	h.Heapify()
	h.PrintHeap()
	fmt.Print(h.Delete())
	h.PrintHeap()

	fmt.Print(h.Delete())
	h.PrintHeap()

	h.Insert(20)
	h.PrintHeap()

	h.Insert(18)
	h.PrintHeap()

	for len(h.data)>0{
		fmt.Println(h.Delete())
	}
}
