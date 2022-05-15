package dp

import "fmt"

var cache []int

func PlayLIS(){
	in := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(in))
}

func lengthOfLIS(nums []int) int {
	cache = make([]int, len(nums))
	maxLen:= 0
	for i:=0;i<len(nums);i++{
		l := lenOfLISHelper(nums, i)
		maxLen = max(maxLen, l)
	}
	return maxLen
}

func lenOfLISHelper(nums []int, start int) int{
	if start == len(nums)-1{
		cache[start] = 1
		return 1
	}
	if cache[start] != 0{
		return cache[start]
	}
	var lis int
	for i:=start+1;i<len(nums);i++{
		if nums[i]>nums[start]{
			lis = max(lis, lenOfLISHelper(nums, i))
		}
	}
	cache[start] = lis+1
	return lis+1

}