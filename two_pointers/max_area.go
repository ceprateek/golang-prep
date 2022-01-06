package two_pointers

import (
	"fmt"
)

/*
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation:
The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
In this case, the max area of water (blue section) the container can contain is 49.
*/

func PlayMaxArea() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxHeight := maxArea(h)
	fmt.Printf("max area: %d", maxHeight)
}

func maxArea(height []int) int {
	var maxArea, left, right, currentArea int
	maxArea = -1
	right = len(height) - 1
	for left < right {
		currentHeight := min(height[right], height[left])
		currentWidth := right - left
		currentArea = currentWidth * currentHeight
		maxArea = max(currentArea, maxArea)
		if height[left]< height[right]{
			left++
		}else if height[left]> height[right]{
			right--
		}else {
			left++
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
