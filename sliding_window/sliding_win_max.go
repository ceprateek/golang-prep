package sliding_window

import (
	"fmt"
)

//https://leetcode.com/problems/sliding-window-maximum/

/*
You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.



Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]
Example 3:

Input: nums = [1,-1], k = 1
Output: [1,-1]
Example 4:

Input: nums = [9,11], k = 2
Output: [11]
Example 5:

Input: nums = [4,-2], k = 2
Output: [4]
*/

func PlayMaxSligingWindow() {
	nums := []int{1, 3, 3, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

type node struct {
	idx int
	val int
}

func maxSlidingWindow(nums []int, k int) []int {
	var out []int
	queue := New()
	for i := 0; i < len(nums); i++ {
		//pop from left if the element is not in window
		if !queue.Empty() {
			left := queue.Left().(node)
			if left.idx < i-k+1 {
				queue.PopLeft()
			}
		}
		currentValToAdd := nums[i]
		for !queue.Empty(){
			right := queue.Right().(node)
			if right.val < currentValToAdd {
				queue.PopRight()
			} else {
				break
			}
		}
		queue.PushRight(node{
			idx: i,
			val: currentValToAdd,
		})
		if i>=k-1{
			out = append(out, queue.Left().(node).val)
		}
	}

	return out
}

func maxOf2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
