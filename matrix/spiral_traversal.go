package matrix

import "fmt"

func PlaySpiral() {
	matrix := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
	interateSpiral(&matrix)
}

func interateSpiral(s *[][]int) {
	matrix := *s
	left := 0
	top := 0
	bottom := len(matrix) - 1
	right := len(matrix[0]) - 1
	direction := 0
	for top<=bottom && left<=right {
		switch direction {
		case 0:
			for i := left; i <= right; i++ {
				fmt.Println(matrix[top][i])
			}
			top++
		case 1:
			for i := top; i <= bottom; i++ {
				fmt.Println(matrix[i][right])
			}
			right--
		case 2:
			for i := right; i >= left; i-- {
				fmt.Println(matrix[bottom][i])
			}
			bottom--

		case 3:
			for i := bottom; i >= top; i-- {
				fmt.Println(matrix[i][left])
			}
			left++
		}
		direction = (direction + 1) % 4
	}
}
