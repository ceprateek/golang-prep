package dp

import "fmt"

func PlayPaths() {
	paths := uniquePaths(3, 7)
	fmt.Println(paths)
}

func uniquePaths(m int, n int) int {
	return uniquePathsHelper(m, n, 0, 0)
}

func uniquePathsHelper(m int, n int, startm, startn int) int {
	if startm == m-1 || startn == n-1 {
		return 1
	}
	rightResult := uniquePathsHelper(m, n, startm+1, startn)
	leftResult := uniquePathsHelper(m, n, startm, startn+1)

	return rightResult + leftResult
}

/*----------------------------------------------*/

//unique paths with obstacles

func PlayPathWithObstacles() {
	obs := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	paths := uniquePathsWithObstaclesDP(obs)
	fmt.Println(paths)

}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return uniquePathsWithObstaclesHepler(obstacleGrid, 0, 0)
}

func uniquePathsWithObstaclesHepler(obstacleGrid [][]int, currentRow, currentCol int) int {
	if currentRow == len(obstacleGrid)-1 && currentCol == len(obstacleGrid[0])-1 {
		return 1
	}
	rightPointValid := isvalidPoint(obstacleGrid, currentRow, currentCol+1)
	downPointValid := isvalidPoint(obstacleGrid, currentRow+1, currentCol)
	var paths int
	if rightPointValid {
		paths = uniquePathsWithObstaclesHepler(obstacleGrid, currentRow, currentCol+1)
	}
	if downPointValid {
		paths += uniquePathsWithObstaclesHepler(obstacleGrid, currentRow+1, currentCol)
	}
	return paths
}

func isvalidPoint(obstacleGrid [][]int, row, col int) bool {
	if row < len(obstacleGrid) && col < len(obstacleGrid[row]) && obstacleGrid[row][col] == 0 {
		return true
	}
	return false
}

/*----------------------------------------------*/

//unique paths with obstacles Dynamic programming

func uniquePathsWithObstaclesDP(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0{
		return 0
	}
	if obstacleGrid[0][0]== 1{
		return 0
	}

	obstacleGrid[0][0] = 1

	for i:=1;i<len(obstacleGrid[0]);i++{
		if obstacleGrid[0][i] == 1{
			obstacleGrid[0][i] = 0
		}else {
			obstacleGrid[0][i] = obstacleGrid[0][i-1]
		}
	}
	for j:=1;j<len(obstacleGrid); j++{
		if obstacleGrid[j][0] == 1{
			obstacleGrid[j][0] = 0
		}else {
			obstacleGrid[j][0] = obstacleGrid[j-1][0]
		}
	}

	for row:=1;row<len(obstacleGrid);row++{
		for col := 1; col<len(obstacleGrid[0]);col++{
			if obstacleGrid[row][col] == 1{
				obstacleGrid[row][col] = 0
			}else {
				obstacleGrid[row][col] = obstacleGrid[row-1][col] + obstacleGrid[row][col-1]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]

}








