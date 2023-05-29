package backtrack

import "fmt"

func PlayLexialOrderLargestSubstring() {
	s := "abab"
	result := largestLexialSubstring(s)
	fmt.Println(result)
}

func largestLexialSubstring(s string) string {
	if len(s) == 0 {
		return s
	}
	largest := rune(s[0])
	largestIndex := 0
	for i, s := range s {
		if s > largest {
			largest = s

		}
	}
	return ""
}
