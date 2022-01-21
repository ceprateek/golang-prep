package greedy

import "fmt"

/*
https://leetcode.com/problems/longest-increasing-subsequence/

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/

func PlayLongestSubsequence() {
	in := []int{10, 9, 2, 5, 3, 7, 101, 18}
	out := lengthOfLIS(in)
	fmt.Println(out)
}


func lengthOfLIS(nums []int) int {
	totalsequeces := make([][]int, 0)
	maxLen := -1
	for i := 0; i < len(nums); i++ {
		element := nums[i]
		merged := false
		for j := 0; j < len(totalsequeces); j++ {
			sq := totalsequeces[j]
			lastElementSq := sq[len(sq)-1]
			if lastElementSq < element{
				newSq := make([]int, len(sq)+1)
				copy(newSq, sq)
				newSq[len(newSq)-1] = element
				totalsequeces = append(totalsequeces, newSq)
				merged = true
				if len(newSq)>maxLen{
					maxLen = len(newSq)
				}
			}
		}
		if !merged{
			totalsequeces = append(totalsequeces, []int{element})
			if maxLen<1{
				maxLen = 1
			}
		}
	}
	return maxLen
}
