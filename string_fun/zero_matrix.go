package string_fun

import "fmt"

func PlayZeroMatrix() {
	matrix := [][]int{
		{1, 1, 1, 1, 1}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 1, 1, 0},
	}
	printArray(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0]= 0
				matrix[0][j]=0
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][0]==0{
				matrix[i][j]= 0
			}
			if matrix[0][j]==0{
				matrix[i][j]= 0
			}
		}
	}
	fmt.Println()
	fmt.Println()
	printArray(matrix)
}

func printArray(in [][]int) {
	fmt.Println()
	for i := 0; i < len(in); i++ {
		fmt.Println(in[i])
	}
}

