package _range

import "strconv"

/*
You are given an inclusive range [lower, upper] and a sorted unique integer array nums,
where all elements are in the inclusive range.
A number x is considered missing if x is in the range [lower, upper] and x is not in nums.

Return the smallest sorted list of ranges that cover every missing number exactly.
That is, no element of nums is in any of the ranges, and each missing number is in
one of the ranges.

Each range [a,b] in the list should be output as:
"a->b" if a != b
"a" if a == b
Example 1:
Input: nums = [0,1,3,50,75], lower = 0, upper = 99
Output: ["2","4->49","51->74","76->99"]
Explanation: The ranges are:
[2,2] --> "2"
[4,49] --> "4->49"
[51,74] --> "51->74"
[76,99] --> "76->99"

Example 2:
Input: nums = [-1], lower = -1, upper = -1
Output: []
Explanation: There are no missing ranges since there are no missing numbers.

*/

func findMissingRanges(nums []int, lower int, upper int) []string {
	out := make([]string, 0)
	l := 0
	r := 1
	if lower < nums[0] {
		s := "[" + strconv.Itoa(lower) + "," + strconv.Itoa(nums[0]-1)
		out = append(out, s)
	}
	for r < len(nums) {
		rl := nums[l] + 1
		rr := nums[r] - 1
		if rr-rl > 0 {
			s := strconv.Itoa(rl) + "->" + strconv.Itoa(rr)
			out = append(out, s)
		} else if rr-rl == 0 {
			s := strconv.Itoa(rl)
			out = append(out, s)
		}
		l++
		r++
	}
	if nums[len(nums)-1] < upper {
		s := strconv.Itoa(nums[len(nums)-1]) + "," + strconv.Itoa(upper)
		out = append(out, s)
	}
	return out
}
