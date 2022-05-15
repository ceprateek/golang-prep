package two_pointers

import "fmt"

func PlayTrap() {
	in := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(in))
}

func trap(height []int) int {
	longestHeightLeft := make([]int, len(height))
	longestHeightRight := make([]int, len(height))
	longestHeightLeft[0] = 0
	for i:=1 ; i< len(height); i++{
		longestHeightLeft[i] = max(longestHeightLeft[i-1], height[i-1])
	}
	longestHeightRight[len(height)-1] = height[len(height)-1]
	for i:=len(height)-2; i>=0;i--{
		longestHeightRight[i] = max(longestHeightRight[i+1], height[i+1])
	}
	var totalwater int
	for i:=1; i<len(height)-1;i++{
		water := min(longestHeightRight[i], longestHeightLeft[i])-height[i]
		if water<0{
			continue
		}
		totalwater += water
	}
	return totalwater
}
