package greedy

import (
	"fmt"
	"strconv"
)

func PlayLargestNumber() {
	s := 20
	d := 3
	val := findLargestNumberWithSumNDigits(s, d)
	fmt.Println(val)
}

func findLargestNumberWithSumNDigits(sum, digits int) int {
	current := ""
	fundLargestHelper(sum, digits, &current)
	val, _ := strconv.Atoi(current)
	return val
}

func fundLargestHelper(sum, digits int, current *string) {
	if sum == 0 {
		return
	} else {
		if sum > 9 {
			*current = *current + "9"
			sum -= 9
		} else {
			*current = *current + strconv.Itoa(sum)
			sum = 0
		}
		fundLargestHelper(sum, digits, current)
	}
}
