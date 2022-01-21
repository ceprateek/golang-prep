package binary_heap

import (
	"fmt"
	"strconv"
)

/*
Given two integer arrays arr1[] and arr2[] sorted in ascending order and an integer k.
Find k pairs with smallest sums such that one element of a pair belongs to arr1[]
and other element belongs to arr2[]
Examples:

Input :  arr1[] = {1, 7, 11}
         arr2[] = {2, 4, 6}
         k = 3
Output : [1, 2],
         [1, 4],
         [1, 6]
Explanation: The first 3 pairs are returned
from the sequence [1, 2], [1, 4], [1, 6],
[7, 2], [7, 4], [11, 2], [7, 6], [11, 4],
[11, 6]
*/

type combination struct {
	idx1, idx2, sum int
}

//create a heap of sums
func findPairsHeap(arr1, arr2 []int, k int) [][]int {
	h := sumHeap{}
	captured := make(map[string]int)
	out := make([][]int, 0, k)
	found := 0
	addToHeap(arr1, arr2, 0, 0, &h, captured)
	for h.Len() > 0 && found < k {
		top := h.Del()
		oneOutput := []int{arr1[top.idx1], arr2[top.idx2]}
		out = append(out, oneOutput)
		found++
		//case i+1,j
		if (top.idx1+1) < len(arr1) && !isConsidered(captured, top.idx1+1, top.idx2) {
			addToHeap(arr1, arr2, top.idx1+1, top.idx2, &h, captured)
		}
		//case i,j+1
		if (top.idx2+1) < len(arr2) && !isConsidered(captured, top.idx1, top.idx2+1) {
			addToHeap(arr1, arr2, top.idx1, top.idx2+1, &h, captured)
		}
	}
	return out
}

func addToHeap(arr1, arr2 []int, i, j int, h *sumHeap, captured map[string]int) {
	s := arr1[i] + arr2[j]
	ci := strconv.Itoa(i) + "," + strconv.Itoa(j)
	captured[ci] = s
	c := combination{
		idx1: i,
		idx2: j,
		sum:  s,
	}
	h.Insert(c)
}

func isConsidered(captured map[string]int, i, j int) bool {
	ci := strconv.Itoa(i) + "," + strconv.Itoa(j)
	if _, ok := captured[ci]; ok {
		return true
	} else {
		return false
	}
}

func PlayFindKPairs() {
	arr1 := []int{1, 7, 11}
	arr2 := []int{2, 4, 6}
	k := 9
	fmt.Println(findPairsHeap(arr1, arr2, k))
}

type sumHeap []combination

func (h *sumHeap) heapifyHelper(root, len int) {
	in := *h
	left := 2*root + 1
	right := 2*root + 2
	min := root

	if left < len && in[left].sum < in[min].sum {
		min = left
	}
	if right < len && in[right].sum < in[min].sum {
		min = right
	}
	if min != root {
		in[min], in[root] = in[root], in[min]
		h.heapifyHelper(min, len)
	}
}

func (h *sumHeap) Heapify() {
	for i := len(*h) / 2; i >= 0; i-- {
		h.heapifyHelper(i, len(*h))
	}
}

func (h *sumHeap) Del() combination {
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	//balance the tree
	h.heapifyHelper(0, len(*h))
	return val
}

func (h *sumHeap) Insert(val combination) {
	*h = append(*h, val)
	h.HeapifyAtIdx(len(*h) - 1)
}

func (h *sumHeap) Len() int {
	return len(*h)
}

func (h *sumHeap) HeapifyAtIdx(idx int) {
	child := idx
	arr := *h
	parent := (child - 1) / 2
	if arr[child].sum < arr[parent].sum {
		arr[child], arr[parent] = arr[parent], arr[child]
		h.HeapifyAtIdx(parent)
	}
}
