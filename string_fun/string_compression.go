package string_fun

import (
	"fmt"
	"strconv"
)

/*
https://leetcode.com/problems/string-compression/

Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
*/

func PlayStringCompression() {
	s := "aabbccc"
	c := []byte(s)
	result := compress(c)
	fmt.Println(result)
}

func compress(chars []byte) int {
	var p int
	result := string(chars[0])
	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[p] {
			if i-p > 1 {
				result = result + strconv.Itoa(i-p)
			}
			result = result + string(chars[i])
			p=i
		}
	}
	if p<len(chars)-1{
		result = result + strconv.Itoa(len(chars)-p)
	}
	fmt.Println(result)
	return len(result)
}
