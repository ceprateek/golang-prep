package greedy

import (
	"fmt"
	"math"
)

func PlayJumpGame() {
	nums := []int{1, 2, 3}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		maxJump := nums[i]
		if maxJump+i >= len(nums)-1 {
			dp[i] = 1
			continue
		}
		minJumps := math.MaxInt32
		for j := 1; j <= maxJump; j++ {
			minJumps = min(minJumps, dp[i+j])
		}
		dp[i] = minJumps+1
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
