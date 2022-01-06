package sliding_window

import "fmt"

/*
Day 4.3: Longest Repeating String after K Replacement. — Medium
*/

/*
We have a string of length n, which consist only UPPER and LOWER CASE characters and we have a number k (always less than n and greater than 0). We can make at most k changes in our string such that we can get a sub-string of maximum length which have all same characters.
Examples:


n = length of string
k = changes you can make
Input : n = 5 k = 2
        str = ABABA
Output : maximum length = 5
which will be (AAAAA)

Input : n = 6 k = 4
       str = HHHHHH
Output : maximum length=6
which will be(HHHHHH)
*/
func Play_longestRepeating(){
	a := "ABABA"
	k := 2
	fmt.Println(longest_repeating_string_after_k_replacements_43(a, k))
}

func longest_repeating_string_after_k_replacements_43(input string, k int) int{
	n := len(input)
	charCounts := make([]int, 26)
	windowStart := 0
	maxLen := 0
	maxCount := 0
	for windowEnd := 0; windowEnd < n; windowEnd++ {
		charCounts[input[windowEnd]-'A']++
		currentCharCount := charCounts[input[windowEnd]-'A']
		//maxCount stores the max number of times a char repeats
		maxCount := max_num(maxCount, currentCharCount)

		//windowEnd - windowStart + 1 - maxCount gives the repeating chars
		//reduce window if the number of chars to change is > k
		for (windowEnd - windowStart + 1 - maxCount) > k {
			charCounts[input[windowEnd]-'A']--
			windowStart++
		}

		maxLen = max_num(maxLen, windowEnd -windowStart+ 1)
	}
	return maxLen
}

func max_num(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
