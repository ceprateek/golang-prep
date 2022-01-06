package sliding_window

import (
	"fmt"
	"math"
)

/*
Day 4.1: Maximum Sum Subarray of size K. — Easy
Day 4.2: Smallest Subarray with a given sum — Easy

*/

func PlaySlidingWindow() {
	//in := []int{3, 2, 1, 4, 5, 1, 1, 1}
	//window := 3
	//maxSumSubarray(&in, window)

	smallsubarr := []int{2, 4, 6, 10, 2, 1}
	//smallestSubarrayWithAGivenSum(smallsubarr, 12)
	smallestSubarrayWithAGivenSumEfficient(smallsubarr, 12)
}

func smallestSubWithSum(arr []int, n, x int) int {
	//  Initialize length of smallest subarray as n+1
	minLen := n + 1

	// Pick every element as starting point
	for start := 0; start < n; start++ {
		// Initialize sum starting with current start
		currSum := arr[start]

		// If first element itself is greater
		if currSum > x {
			return 1
		}

		// Try different ending points for curremt start
		for end := start + 1; end < n; end++ {
			// add last element to current sum
			currSum += arr[end]

			// If sum becomes more than x and length of
			// this subarray is smaller than current smallest
			// length, update the smallest length (or result)
			if currSum > x && (end-start+1 < minLen) {
				minLen = end - start + 1
			}

		}
	}
	return minLen
}

func maxSumSubarray(in *[]int, window int) {
	arr := *in
	N := len(arr)
	var winSum, max int

	for i := 0; i < window; i++ {
		winSum += arr[i]
	}
	max = winSum
	for i := 1; i <= N-window; i++ {
		winSum = winSum - arr[i-1] + arr[i+window-1]
		max = maxOfInts(winSum, max)
	}
	fmt.Println(max)
}

func maxOfInts(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func smallestSubarrayWithAGivenSum(in []int, sum int) {
	var tempsum, tempmin, min int
	min = math.MaxInt32
	for i := 0; i < len(in); i++ {
		tempmin = 0
		tempsum = 0

		for j := i; j < len(in); j++ {
			tempmin++
			tempsum += in[j]
			if tempsum == sum {
				min = minNum(min, tempmin)
				break
			} else if tempsum > sum {
				break
			}
		}
	}
	fmt.Println(min)
}

func smallestSubarrayWithAGivenSumEfficient(in []int, sum int) {
	var left, right, currentSum, minWindow, start int
	minWindow = math.MaxInt32
	currentSum = in[0]
	for left < len(in) && right < len(in) {
		if left == right {
			if currentSum == sum {
				minWindow = 1
				start = left
				fmt.Printf("window:%d start:%d\n", minWindow, start)
			} else {
				right++
				if right >= len(in) {
					break
				}
				currentSum += in[right]
			}
		}
		if currentSum < sum {
			right++
			if right >= len(in) {
				break
			}
			currentSum += in[right]
		} else {
			if currentSum == sum {
				start = left
				fmt.Printf("window:%d start:%d\n", minWindow, start)
				minWindow = minNum(minWindow, (right-left+1))
				currentSum -= in[left]
				left++
			}else {
				currentSum -= in[left]
				left++
			}
		}
	}
	fmt.Printf("window:%d start:%d", minWindow, start)
}

func minNum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
