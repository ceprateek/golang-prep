package two_pointers

import "fmt"

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array.
Example 1:
Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

Constraints:
1 <= nums.length <= 105
nums[i] is either 0 or 1.
*/

func PlayMaxConsecutiveOnes() {
	in := []int{1}
	fmt.Println(findMaxConsecutiveOnes(in))
}

func findMaxConsecutiveOnes(nums []int) int {
	var maxWindow int
	for i := 0; i < len(nums); {
		if nums[i] != 1{
			i++
			continue
		}
		window := findMaxPerIndex(nums, i)
		i += window
		maxWindow = max(maxWindow, window)
	}
	return maxWindow
}

func findMaxPerIndex(nums []int, i int) int {
	currentWindow := 1
	for j:=i+1;j<len(nums);j++{
		if nums[j] != 1{
			return currentWindow
		}else {
			currentWindow++
		}
	}
	return currentWindow
}
