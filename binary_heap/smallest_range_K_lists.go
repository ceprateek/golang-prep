package binary_heap

import (
	"fmt"
	"math"
	"sort"
)

/*
Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
Output: [20,24]
Explanation:
List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
List 2: [0, 9, 12, 20], 20 is in range [20,24].
List 3: [5, 18, 22, 30], 22 is in range [20,24].
*/

func smallestRange(nums [][]int) []int {
	result := make([]int, 2)
	minRange := math.MaxInt32
	idsInRange := make([]int, len(nums))
	var merged []*mergedListVal
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			m := &mergedListVal{nums[i][j], i}
			merged = append(merged, m)
		}
	}
	sort.Slice(merged, func(i, j int) bool {
		if merged[i].val < merged[j].val {
			return true
		} else {
			return false
		}
	})
	var total, j int
	for i := 0; i < len(merged); i++ {
		idsInRange[merged[i].idx]++
		if idsInRange[merged[i].idx] == 1 {
			total++
		}
		if total == len(nums) {
			//range found
			rightOfRange := merged[i].idx
			j = 0
			for j<i{
				left := merged[j].idx
				if idsInRange[left] > 1 {
					idsInRange[left]--
					j++
				} else {
					break
				}
			}
			leftOfRange := j
			rangeVal := merged[rightOfRange].val - merged[leftOfRange].val
			if rangeVal < minRange {
				result[0] = merged[leftOfRange].val
				result[1] = merged[rightOfRange].val
				minRange = rangeVal
			}
		}
	}
	return result
}

func PlaySmallestRange() {
	in := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	out := smallestRange(in)
	fmt.Println(out)
}

type mergedListVal struct {
	val int
	idx int
}
