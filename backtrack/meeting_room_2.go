package backtrack

import (
	"fmt"
	"sort"
)

//intervals = [[0,30],[5,10],[15,20]]

func PlayMinMeetingRooms() {
	intervals := [][]int{{1, 3}, {2, 6}, {4, 30}, {5, 10}, {15, 20}}
	rooms := minMeetingRooms(intervals)
	fmt.Println(rooms)
}

func minMeetingRooms(intervals [][]int) int {
	meetings := make([]*meeting, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		meet := &meeting{
			start: interval[0],
			end:   interval[1],
		}
		meetings = append(meetings, meet)
	}
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].start < meetings[j].start {
			return true
		}
		return false
	})
	m := minHeap{}
	m.Insert(meetings[0])
	for i := 1; i < len(meetings); i++ {
		meet := meetings[i]
		room := m.Pop()
		if room.end <= meet.start {
			m.Insert(meet)
		} else {
			m.Insert(meet)
			m.Insert(room)
		}
	}
	return len(m)
}

type meeting struct {
	start, end int
}

type minHeap []*meeting

func (m *minHeap) Insert(meet *meeting) {
	*m = append(*m, meet)
	m.heapifyAtIdx(len(*m) - 1)
}

func (m *minHeap) heapifyAtIdx(idx int) {
	parent := (idx - 1) / 2
	if (*m)[idx].end < (*m)[parent].end {
		(*m)[idx], (*m)[parent] = (*m)[parent], (*m)[idx]
		m.heapifyAtIdx(parent)
	}
}

func (m *minHeap) Pop() *meeting {
	(*m)[0], (*m)[len(*m)-1] = (*m)[len(*m)-1], (*m)[0]
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	m.heapifyHelper(0, len(*m))
	return val
}

func (h *minHeap) heapifyHelper(root, len int) {
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

/*
child = parent*2 +1
child = parent*2 +2

parent = (child-1)/2
*/
