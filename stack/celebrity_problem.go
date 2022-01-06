package stacks

import "fmt"

/*
MATRIX = { {0, 0, 1, 0},
           {0, 0, 1, 0},
           {0, 0, 0, 0},
           {0, 0, 1, 0} }
Output:id = 2
*/

func PlayCelebrity() {
	in := [][]int{
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
	}
	playCelebrityWork(in)
}

func playCelebrityWork(input [][]int) {
	isRegular := make([]bool, len(input))
	for i := 0; i < len(input); i++ {
		if isRegular[i] {
			continue
		}
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 1 {
				isRegular[i] = true
			} else if i != j {
				isRegular[j] = true
			}
		}
	}
	for i, val := range isRegular {
		if !val {
			fmt.Println(i)
		}
	}
}
