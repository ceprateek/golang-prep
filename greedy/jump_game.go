package greedy

import "fmt"

/*
https://leetcode.com/problems/jump-game/

Input: nums = [2,3,1,1,4]
true
Input: nums = [3,2,1,0,4]
false

note: this is not a greedy solution
however a solution which is exploring everything.
I am leaving greedy as a TODO

*/

func PlayJumper()  {
	//in := []int{2,3,1,1,4}
	in := []int{3,2,1,0,4}
	fmt.Println(canJump(in))
}

func canJump(nums []int) bool {
	return canJumpHelper(nums, 0)
}

func canJumpHelper(nums []int, current int)bool{
	if current == len(nums)-1{
		return true
	}
	for i:=1;i<=nums[current];i++{
		if (current+i)<len(nums){
			if canJumpHelper(nums, current+i){
				return true
			}
		}
	}
	return false
}
