package matrix

import "fmt"

func rotate_inplace(in *[][]int) {
	mat := *in
	N := len(mat)
	for x := 0; x < N/2; x++ {
		// Consider elements in group
		// of 4 in current square
		for y := x; y < N-x-1; y++ {
			// Store current cell in
			// temp variable
			temp := mat[x][y]

			// Move values from right to top
			mat[x][y] = mat[y][N-1-x]

			// Move values from bottom to right
			mat[y][N-1-x] = mat[N-1-x][N-1-y]

			// Move values from left to bottom
			mat[N-1-x][N-1-y] = mat[N-1-y][x]

			// Assign temp to left
			mat[N-1-y][x] = temp
		}
	}
}

func PlayRotateMatrix() {
	in := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	rotate_inplace(&in)
	for _, val := range in {
		fmt.Println(val)
	}
}
