package greedy

import (
	"fmt"
	"sort"
)

/*
https://leetcode.com/problems/reorganize-string/
Input: s = "aab"
Output: "aba"

Input: s = "aaab"
Output: ""
*/

func PlayReorganizeString() {
	s := "abbabbaaab"
	o := reorganizeString(s)
	fmt.Println(o)
}

func reorganizeString(s string) string {
	heap := priorityQueueMaxFreq{}
	chars := make(map[rune]int)
	for _, c := range s {
		if _, ok := chars[c]; ok {
			chars[c]++
		} else {
			chars[c] = 1
		}
	}
	for c, f := range chars {
		cf := charWithFreq{
			char: c,
			freq: f,
		}
		heap.Insert(&cf)
	}
	var out string
	for len(heap) > 0 {
		mostFrequent := heap.Delete()
		out = out + string(mostFrequent.char)
		mostFrequent.freq--
		if mostFrequent.freq == 0 {
			continue
		}
		if heap.Len() == 0 {
			return "not possible"
		}
		secondFrequent := heap.Delete()
		out = out + string(secondFrequent.char)
		secondFrequent.freq--
		heap.Insert(mostFrequent)
		if secondFrequent.freq>0{
			heap.Insert(secondFrequent)
		}

	}
	return out
}

type charWithFreq struct {
	char rune
	freq int
}

type priorityQueueMaxFreq []*charWithFreq

func (p *priorityQueueMaxFreq) Len() int {
	return len(*p)
}

func (p *priorityQueueMaxFreq) Insert(val *charWithFreq) {
	*p = append(*p, val)
	sort.Slice(*p, func(i, j int) bool {
		if (*p)[i].freq > (*p)[j].freq {
			return true
		} else {
			return false
		}
	})
}

func (p *priorityQueueMaxFreq) Delete() *charWithFreq {
	if len(*p) == 0 {
		return nil
	}
	val := (*p)[0]
	*p = (*p)[1:]
	if len(*p) > 1 {
		sort.Slice(*p, func(i, j int) bool {
			if (*p)[i].freq > (*p)[j].freq {
				return true
			} else {
				return false
			}
		})
	}
	return val
}
