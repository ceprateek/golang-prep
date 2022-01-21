package greedy

import (
	"fmt"
	"sort"
)

/*
https://leetcode.com/problems/maximum-profit-in-job-scheduling/
Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
Output: 120
Explanation: The subset chosen is the first and fourth job.
Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.
*/

func jobScheduling(startTime, endTime, profit []int) int {
	jobs := make([]*Job, 0, len(startTime))
	for i := 0; i < len(startTime); i++ {
		j := &Job{startTime[i], endTime[i], profit[i]}
		jobs = append(jobs, j)
	}
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].start<jobs[j].start{
			return true
		}
		return false
	})
	maxProfit := 0
	jobChains := make([]*Job, 0)
	for i := 0; i < len(jobs); i++ {
		merged := false
		for j:=0;j<len(jobChains);j++{
			jobChain := jobChains[j]
			if jobChain.end<=jobs[i].start{
				newJobChain := &Job{}
				newJobChain.start = jobChain.start
				newJobChain.end = jobs[i].end
				chainProfit := jobChain.profit + jobs[i].profit
				newJobChain.profit = chainProfit
				if chainProfit>maxProfit{
					maxProfit = chainProfit
				}
				jobChains = append(jobChains, newJobChain)
				merged = true
			}
		}
		if !merged{
			jobChains = append(jobChains, jobs[i])
			if jobs[i].profit>maxProfit{
				maxProfit = jobs[i].profit
			}
		}
	}

	return maxProfit
}

type Job struct {
	start, end, profit int
}

func PlayJobSequenceLeet() {
	start :=  []int{1, 2, 3, 3}
	end :=    []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}
	totalprofit := jobScheduling(start, end, profit)
	fmt.Println(totalprofit)
}
