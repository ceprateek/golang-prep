package greedy

import (
	"fmt"
	"sort"
)

/*
https://www.geeksforgeeks.org/job-sequencing-problem/

Input: Four Jobs with following
deadlines and profits
JobID  Deadline  Profit
  a      4        20
  b      1        10
  c      1        40
  d      1        30
Output: Following is maximum
profit sequence of jobs
        c, a

*/

type job struct {
	id               string
	deadline, profit int
}

func PlayJobSequence() {
	j1 := &job{"a", 2, 100}
	j2 := &job{"b", 1, 19}
	j3 := &job{"c", 2, 27}
	j4 := &job{"d", 1, 25}
	j5 := &job{"e", 3, 15}
	in := []*job{j1, j2, j3, j4, j5}
	fmt.Println(jobSequence(in))
}

func jobSequence(in []*job) []string {
	out := make([]string, len(in))
	sort.Slice(in, func(i, j int) bool {
		if in[i].profit > in[j].profit {
			return true
		}
		return false
	})
	slots := make([]bool, len(in))
	for i := 0; i < len(in); i++ {
		maxSlot := in[i].deadline-1
		for j:=maxSlot;j>=0;j--{
			if !slots[j]{
				slots[j] = true
				out[j] = in[i].id
				break
			}
		}
	}

	return out
}
