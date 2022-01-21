package binary_heap

import (
	"sort"
)

/*
Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
Output: [[3,4]]
Explanation: There are a total of three employees, and all common
free time intervals would be [-inf, 1], [3, 4], [10, inf].
We discard any intervals that contain inf as they aren't finite.
Example 2:

Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
Output: [[5,6],[7,9]]
*/

func PlayFreeTime() {

}

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	heap := MaxHeapInterval{}
	result := make([]*Interval, 0)
	intervals := make([]*Interval, 0)
	for i := 0; i < len(schedule); i++ {
		for j := 0; j < len(schedule[i]); j++ {
			m := &Interval{
				Start: schedule[i][j].Start,
				End:   schedule[i][j].End,
			}
			intervals = append(intervals, m)
		}
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		} else {
			return false
		}
	})
	heap.Insert(intervals[0])
	for i := 1; i < len(intervals); i++ {
		top := heap.Del()
		interval := intervals[i]
		if interval.Start >= top.Start && interval.Start <= top.End {
			top.End = max(interval.End, top.End)
			heap.Insert(top)
		} else {
			heap.Insert(top)
			heap.Insert(interval)
		}
	}
	if len(heap) <= 1 {
		return nil
	}
	previous := heap.Del()
	for len(heap) > 0 {
		m := heap.Del()
		i := Interval{
			Start: m.End,
			End:   previous.Start,
		}
		previous = m
		result = append(result, &i)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MaxHeapInterval []*Interval

func (h *MaxHeapInterval) heapifyHelper(root, len int) {
	in := *h
	left := 2*root + 1
	right := 2*root + 2
	max := root

	if left < len && in[left].End > in[max].End {
		max = left
	}
	if right < len && in[right].End > in[max].End {
		max = right
	}
	if max != root {
		in[max], in[root] = in[root], in[max]
		h.heapifyHelper(max, len)
	}
}

func (h *MaxHeapInterval) Del() *Interval {
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	//balance the tree
	h.heapifyHelper(0, len(*h))
	return val
}

func (h *MaxHeapInterval) Insert(val *Interval) {
	*h = append(*h, val)
	h.HeapifyAtIdx(len(*h) - 1)
}

func (h *MaxHeapInterval) HeapifyAtIdx(idx int) {
	child := idx
	arr := *h
	parent := (child - 1) / 2
	if arr[child].End > arr[parent].End {
		arr[child], arr[parent] = arr[parent], arr[child]
		h.HeapifyAtIdx(parent)
	}
}
