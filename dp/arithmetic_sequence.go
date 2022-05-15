package dp

import "fmt"

func PlayArithmeticSequence(){
	fmt.Println(numberOfArithmeticSlices([]int{1}))
}

var sum int

func numberOfArithmeticSlices(nums []int) int {
	numberOfArithemticSlicesHelper(nums, len(nums)-1)
	return sum
}


func numberOfArithemticSlicesHelper(nums []int, end int) int{
	if end<2{
		return 0
	}
	s := 0
	if nums[end]-nums[end-1] == nums[end-1]-nums[end-2]{
		s = 1 + numberOfArithemticSlicesHelper(nums, end-1)
		sum += s
	}else{
		numberOfArithemticSlicesHelper(nums, end-1)
	}
	return s
}
