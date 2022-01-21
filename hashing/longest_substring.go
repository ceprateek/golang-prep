package hashing

import (
	"fmt"
	"strings"
)

/*
longest substring without repeating characters
For “ABDEFGABEF”, the longest substring are “BDEFGA” and “DEFGAB”, with length 6.
For “BBBB” the longest substring is “B”, with length 1.
For “GEEKSFORGEEKS”, there are two longest substrings shown in the below diagrams, with length 7
*/

func findLongestSubstring(input string) int {
	var left, right, window int
	if len(input) == 0 {
		return 0
	} else if len(input) == 1 {
		return 1
	}
	right = 1
	for right < len(input) {
		char := input[right]
		if strings.Contains(input[left:right], string(char)) {
			left = strings.LastIndex(input[left:right], string(char)) + 1+left
		}
		window = max(window, right-left+1)
		right++
		fmt.Printf("left: %d right:%d window: %s\n", left, right, input[left:right])
	}
	return window
}

func PlayLongestSubstring() {
	input := "GEEKSFORGEEKS"
	window := findLongestSubstring(input)
	fmt.Println(window)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
