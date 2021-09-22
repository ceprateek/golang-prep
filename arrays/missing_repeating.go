package arrays

import (
	"fmt"
	"sort"
	"strconv"
)

func PlayRepeating()  {
	findMissingRepeatingNumber([]int{1,2,3,4,5,5})
}

func findMissingRepeatingNumber(in []int) {
	sort.Ints(in)
	for i := 1; i < len(in); i++ {
		fmt.Println(fmt.Sprintf("i: %d in[i]: %d in[i-1]: %d", i, in[i], in[i-1]))
		if in[i-1] == in[i] {
			fmt.Println("repeating: " + strconv.Itoa(in[i]))
			fmt.Println("missing: " + strconv.Itoa(in[i]+1))
		}
	}
}
