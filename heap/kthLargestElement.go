package heap_internal

import (
	"container/heap"
	"fmt"
)

type IMinHeapSizeK interface {
	GetSize() int
	Insert(n int)
	GetSmallestElement() int
}

type MinHeapSizeK struct {
	maxSize int
	data    []int
}

func (m *MinHeapSizeK) Insert(n int) {
	if m.GetSize() == m.maxSize && n < m.GetSmallestElement() {
		return
	}
	m.insert(n)

	if m.GetSize() > m.maxSize {
		m.delete()
	}
}

func (m *MinHeapSizeK) insert(n int) {
	m.data = append(m.data, n)
	m.heapifyInsert(len(m.data) - 1)
}

func (m *MinHeapSizeK) heapifyInsert(idx int) {
	if idx < 0 {
		return
	}
	parentIdx := (idx - 1) / 2
	if m.data[parentIdx] > m.data[idx] {
		m.data[parentIdx], m.data[idx] = m.data[idx], m.data[parentIdx]
		m.heapifyInsert(parentIdx)
	}

}

func (m *MinHeapSizeK) delete() {
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.heapifyDelete(0)
}

func (m *MinHeapSizeK) heapifyDelete(idx int) {
	smallest := idx
	left := idx*2 + 1
	right := idx*2 + 2

	if left < len(m.data) && m.data[left] < m.data[smallest] {
		smallest = left
	}
	if right < len(m.data) && m.data[right] < m.data[smallest] {
		smallest = right
	}
	if smallest != idx {
		m.data[idx], m.data[smallest] = m.data[smallest], m.data[idx]
		m.heapifyDelete(smallest)
	}

}

func (m *MinHeapSizeK) GetSize() int {
	return len(m.data)
}

func (m *MinHeapSizeK) GetSmallestElement() int {
	return m.data[0]
}

func NewMinHeapSizeK(k int) IMinHeapSizeK {
	m := MinHeapSizeK{
		data:    make([]int, 0),
		maxSize: k,
	}
	return &m
}

func PlayKthElement() {
	input := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	// myHeap := NewMinHeapSizeK(k)
	// for _, i := range input {
	// 	myHeap.Insert(i)
	// }
	// fmt.Println(myHeap.GetSmallestElement())
	t := &testHeap{}
	heap.Init(t)
	for _, val := range input {
		heap.Push(t, val)
		if t.Len() > k {
			heap.Pop(t)
		}
		fmt.Println(*t)
	}
	fmt.Print(heap.Pop(t))
}

/////////////////////////////////////////////////////////////

type testHeap []int

func (t *testHeap) Len() int {
	return len(*t)
}

func (t *testHeap) Less(i, j int) bool {
	return (*t)[i] < (*t)[j]
}

func (t *testHeap) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *testHeap) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *testHeap) Pop() interface{} {
	last := (*t)[t.Len()-1]
	*t = (*t)[:t.Len()-1]
	return last
}
