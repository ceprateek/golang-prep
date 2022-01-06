package sliding_window

import "fmt"


/*
the best approach is to increase the window in the loop with right always expanding
then check for the negative condition of the window in a for loop until the negative condition is negated
 */
func maxWindowNoRepeatChars(in string) {
	var left, currentMax, maxLen int
	var currentWindow string
	charsCurrentWindow := make([]int, 128)
	for right := 0; right < len(in); right++ {
		currentWindow = currentWindow + string(in[right])
		charsCurrentWindow[in[right]]++
		currentMax++
		for charsCurrentWindow[in[right]] > 1 {
			elementToRemove := currentWindow[0]
			currentWindow = currentWindow[1:]
			left++
			charsCurrentWindow[elementToRemove]--
			currentMax--
		}
		maxLen = maxOfInts(maxLen, currentMax)
	}
	fmt.Println(maxLen)
}

func PlayMaxSubNonRepeating()  {
	s := "abcabcdbb"
	maxWindowNoRepeatChars(s)
}
