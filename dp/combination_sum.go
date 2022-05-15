package dp

import "fmt"

func PlayCombinationSun4(){
	nums := []int{9}
	fmt.Println(combinationSum4(nums, 3))
}

var combinations int

func combinationSum4(nums []int, target int) int {
	combinationSumHelper(nums, target, 0, &[]int{})
	return combinations
}

func combinationSumHelper(nums []int, target, currentSum int, combination *[]int) {
	if currentSum > target {
		return
	}
	if currentSum == target {
		combinations++
		return
	}

	for i := 0; i < len(nums); i++ {
		chosen := nums[i]
		currentSum += chosen
		if currentSum <= target {
			*combination = append(*combination, chosen)
			combinationSumHelper(nums, target, currentSum, combination)
			*combination = (*combination)[:len(*combination)-1]
		}
		currentSum -= chosen
	}
}