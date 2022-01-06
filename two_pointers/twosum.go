package two_pointers

import (
	"fmt"
	"sort"
)

/*
Input: arr[] = {0, -1, 2, -3, 1}
        sum = -2
Output: -3, 1
*/

func PlayTwoSum() {
	a := []int{0, -1, -2, -3, 1}
	twoSum(a, -2)
}

func twoSum(nums []int, target int) {
	sort.Ints(nums)
	fmt.Println(nums)
	i := 0
	j := len(nums) - 1
	for i < j {
		sum := nums[i] + nums[j]
		if sum < target{
			i++
		} else if sum > target {
			j--
		} else {
			fmt.Printf("found match: %d %d\n", nums[i], nums[j])
			i++
			j--
		}
	}
}
