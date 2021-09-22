package backtrack

import (
	"fmt"
)

var total int

func Permutate(input string) {
	permutateHelper(input, "")
	fmt.Println(total)
}

func indent(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("   ")
	}
}

func permutateHelper(input string, temp string) {
	indent(len(temp))
	fmt.Println(fmt.Sprintf("input: %s temp: %s", input, temp))
	if len(input) == 0 {
		fmt.Println(temp)
		total++
	} else {
		for i := 0; i < len(input); i++ {
			//choose
			c := input[i]
			indent(len(temp))
			fmt.Println("chosen: " + string(c))
			//explore
			tempInput := input
			input = input[0:i] + input[i+1:]
			temp = temp + string(c)
			permutateHelper(input, temp)

			//unchoose
			input = tempInput
			temp = temp[:len(temp)-1]
			indent(len(temp))
			fmt.Println(fmt.Sprintf("unchoose c: %s input: %s temp: %s", string(c), input, temp))
		}
	}
}
