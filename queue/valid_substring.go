package queue

import (
	"fmt"

	stacks "github.com/ceprateek/golang-prep/stack"
)

/*
Input : ((()
Output : 2
Explanation : ()

Input: )()())
Output : 4
Explanation: ()()

Input:  ()(()))))
Output: 6
Explanation:  ()(())
*/

func findLongestValidSubstring(in string) int {
	chars := []rune(in)
	stack := stacks.Stack{}
	var maxLen int
	for i := 0; i < len(chars); i++ {
		val := string(chars[i])
		switch val {
		case "(":
			stack.Push(i)
		case ")":
			if stack.IsEmpty() {
				stack.Clear()
			} else {
				idx := stack.Pop().(int)
				maxLen = max(maxLen, i-idx+1)
			}
		default:
			panic("wrong string")

		}
		if string(chars[i]) == "(" {

		} else if string(chars[i]) == ")" {

		}
	}
	return maxLen
}

func PlayFindLongestSubstring() {
	out := findLongestValidSubstring("(()())(")
	fmt.Println(out)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
