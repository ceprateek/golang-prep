package recursion

import (
	"fmt"
)

/*
There is only one character 'A' on the screen of a notepad.
You can perform two operations on this notepad for each step:

Copy All: You can copy all the characters present on the screen
Paste: You can paste the characters which are copied last time.

Input: n = 3
Output: 3
Explanation: Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.

https://leetcode.com/problems/2-keys-keyboard/
*/
func PlaySpecialKeyboard(){
	n := 7
	fmt.Println(minSteps(n))
}

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	var minStep = n
	for i := 2; i <= n; i++ {
		if n % i == 0 {
			tmpMin := minSteps(i) + n / i
			if tmpMin < minStep {
				minStep = tmpMin
			}
		}
	}
	return minStep
}