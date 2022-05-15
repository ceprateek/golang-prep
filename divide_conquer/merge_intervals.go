package divide_conquer

import (
	"fmt"
	"sort"
)

func PlayMergeIntervals() {
	intervals := [][]int{{1, 3}, {1, 5}, {2, 6}, {8, 10}, {15, 18}}
	k := merge(intervals)
	fmt.Println(k)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return false
		}
	})
	var merged [][]int
	merged = append(merged, intervals[0])
	for k := 1; k < len(intervals); k++ {
		lastMerged := merged[len(merged)-1]
		toMerge := intervals[k]
		if toMerge[0]>=lastMerged[0] && toMerge[0]<=lastMerged[1] && toMerge[1]>lastMerged[1]{
			lastMerged[1] = toMerge[1]
			merged = merged[:len(merged)-1]
			merged = append(merged, lastMerged)
		}
		if toMerge[0]>lastMerged[1]{
			merged = append(merged, toMerge)
		}
	}
	return merged
}
