package hashing

import (
	"fmt"
)

/*
Input: arr[] = {15, -2, 2, -8, 1, 7, 10, 23};
Output: 5
Explanation: The longest sub-array with
elements summing up-to 0 is {-2, 2, -8, 1, 7}

Input: arr[] = {1, 2, 3}
Output: 0
Explanation:There is no subarray with 0 sum

Input:  arr[] = {1, 0, 3}
Output:  1
Explanation: The longest sub-array with
elements summing up-to 0 is {0}
*/

func largestSubarraySumZero(nums []int) int {
	//find sums at each index
	//if the same sum re-occurs, that means that all elements in between are summing to zero
	cache := make(map[int]int, len(nums))
	cache[-1] = 0
	currentSum := 0
	window := -1
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		if idx, ok := cache[currentSum]; ok {
			window = max(window, i-idx)
			fmt.Printf("%d %d %d\n", idx+1, i, window)
		}else {
			cache[currentSum] = i
		}
	}
	return window
}

func PlayLargestSubarraySum() {
	arr := []int{2, 8, -3, -5, 2, -4, 6, 1, 2, 1, -3, 4}
	out := largestSubarraySumZero(arr)
	fmt.Println(out)
}
