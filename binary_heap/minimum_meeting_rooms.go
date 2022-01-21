package binary_heap

import (
	"fmt"
	"sort"
)

/*
Given an array of meeting time intervals where intervals[i] = [starti, endi],
return the minimum number of conference rooms required.

Example 1:
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2
Example 2:

Input: intervals = [[7,10],[2,4]]
Output: 1
*/

func minMeetingRooms(intervals [][]int) int {
	in := make([]*meeting, 0, len(intervals))
	for _, meet := range intervals {
		m := meeting{meet[0], meet[1]}
		in = append(in, &m)
	}
	sort.Slice(in, func(i, j int) bool {
		if in[i].start < in[j].start {
			return true
		} else {
			return false
		}
	})
	var heap minHeapMeeting
	heap.Insert(in[0])
	for i := 1; i < len(in); i++ {
		m := in[i]
		room := heap.Del()
		if room.end > m.start {
			heap.Insert(room)
			heap.Insert(m)
		}else {
			heap.Insert(m)
		}

	}
	return len(heap)
}

type meeting struct {
	start, end int
}

func PlayMinMeetingRooms() {
	meetings := [][]int{{0, 30}, {5, 10}, {15, 20}}
	out := minMeetingRooms(meetings)
	fmt.Println(out)
}

type minHeapMeeting []*meeting

func (h *minHeapMeeting) heapifyHelper(root, len int) {
	in := *h
	left := 2*root + 1
	right := 2*root + 2
	min := root

	if left < len && in[left].end < in[min].end {
		min = left
	}
	if right < len && in[right].end < in[min].end {
		min = right
	}
	if min != root {
		in[min], in[root] = in[root], in[min]
		h.heapifyHelper(min, len)
	}
}

func (h *minHeapMeeting) Heapify() {
	for i := len(*h) / 2; i >= 0; i-- {
		h.heapifyHelper(i, len(*h))
	}
}

func (h *minHeapMeeting) Del() *meeting {
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	val := (*h)[len((*h))-1]
	*h = (*h)[:len(*h)-1]
	//balance the tree
	h.heapifyHelper(0, len(*h))
	return val
}

func (h *minHeapMeeting) Insert(val *meeting) {
	*h = append(*h, val)
	h.HeapifyAtIdx(len(*h) - 1)
}

func (h *minHeapMeeting) HeapifyAtIdx(idx int) {
	child := idx
	arr := *h
	parent := (child - 1) / 2
	if arr[child].end < arr[parent].end {
		arr[child], arr[parent] = arr[parent], arr[child]
		h.HeapifyAtIdx(parent)
	}
}
