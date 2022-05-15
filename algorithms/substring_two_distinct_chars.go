package algorithms

import "math"

/*
Given a string s, return the length of the longest substring that
contains at most two distinct characters.

Input: s = "ccaabbb"
Output: 5
Explanation: The substring is "aabbb" which its length is 5.

*/

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) < 3 {
		return len(s)
	}
	var maxLen,left, right int
	cache := make(map[uint8]int)

	for right < len(s) {
		cache[s[right]] = right
		right++
		if len(cache) == 3 {
			delIdx := findMinValMap(cache)
			left = delIdx+1
			delete(cache, s[delIdx])
		}
		maxLen = int(math.Max(float64(maxLen), float64(right-left+1)))
	}
	return maxLen
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) <= k{
		return len(s)
	}
	var maxLen,left, right int
	cache := make(map[uint8]int)

	for right < len(s) {
		cache[s[right]] = right
		right++
		if len(cache) == k+1 {
			delIdx := findMinValMap(cache)
			left = delIdx+1
			delete(cache, s[delIdx])
		}
		maxLen = int(math.Max(float64(maxLen), float64(right-left+1)))
	}
	return maxLen
}

func findMinValMap(cache map[uint8]int) int {
	min := math.MaxInt32
	for _,val := range cache{
		if val<min{
			min = val
		}
	}
	return min
}
