package arrays

import "fmt"

func PlaySetZerosMatrix() {
	in := [][]int{{1, 1, 1, 1, 1}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 1, 1, 0}}
	setZerosInMatrix(&in)
	printArray(in)
}

type point struct {
	row, col int
}

func setZerosInMatrix(in *[][]int) {
	input := *in
	var zeros []point
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				zero := point{
					row: i,
					col: j,
				}
				zeros = append(zeros, zero)
			}
		}
	}
	for _, zero := range zeros {
		setRowColZero(zero.row, zero.col, in)
	}
}

func setRowColZero(row, col int, in *[][]int) {
	input := *in
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if i == row {
				input[i][j] = 0
			}
			if j == col {
				input[i][j] = 0
			}
		}
	}
}

func printArray(in [][]int) {
	fmt.Println()
	for i := 0; i < len(in); i++ {
		fmt.Println(in[i])
	}
}
