package two_pointers

import "fmt"

func PlayTrap() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	water := trap(height)
	fmt.Println(water)
}

func trap(height []int) int {
	var result int
	var left,right []int
	left = make([]int, len(height))
	right = make([]int, len(height))
	left[0] = height[0]
	for i:=1;i<len(height);i++{
		left[i] = max(left[i-1], height[i])
	}
	right[len(height)-1] = height[len(height)-1]
	for i:=len(height)-2;i>=0;i--{
		right[i] = max(right[i+1], height[i])
	}
	for i := 0; i < len(height); i++ {
		newWater := min(left[i], right[i])-height[i]
		result+=newWater
	}
	return result
}
