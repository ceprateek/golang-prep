package binary_heap

type MinHeap []int

func (h *MinHeap) heapifyHelper(root, len int) {
	in := *h
	left := 2*root + 1
	right := 2*root + 2
	min := root

	if left < len && in[left] < in[min] {
		min = left
	}
	if right < len && in[right] < in[min] {
		min = right
	}
	if min != root {
		in[min], in[root] = in[root], in[min]
		h.heapifyHelper(min, len)
	}
}

func (h *MinHeap) Heapify() {
	for i := len(*h) / 2; i >= 0; i-- {
		h.heapifyHelper(i, len(*h))
	}
}

func (h *MinHeap) Del() int {
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	val := (*h)[len((*h))-1]
	*h = (*h)[:len(*h)-1]
	//balance the tree
	h.heapifyHelper(0, len(*h))
	return val
}

func (h *MinHeap) Insert(val int) {
	*h = append(*h, val)
	h.HeapifyAtIdx(len(*h) - 1)
}

func (h *MinHeap) HeapifyAtIdx(idx int){
	child := idx
	arr := *h
	parent := (child-1)/2
	if arr[child] > arr[parent]{
		arr[child], arr[parent] = arr[parent], arr[child]
		h.HeapifyAtIdx(parent)
	}
}