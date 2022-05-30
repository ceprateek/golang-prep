package recursion

import (
	"fmt"
	"math"
)

func PlayShortestPath() {
	g := [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}
	r := shortestPath(g, 1)
	fmt.Println(r)
}

func shortestPath(grid [][]int, k int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}
	visited[0][0] = true
	return shortestPathHelper(grid, &visited, k, []int{0, 0}, 0)
}

var rows = []int{0, 0, -1, 1}
var cols = []int{-1, 1, 0, 0}

func shortestPathHelper(grid [][]int, visited *[][]bool, k int, current []int, steps int) int {
	if current[0] == len(grid)-1 && current[1] == len(grid[current[0]])-1 {
		return steps
	}
	if grid[current[0]][current[1]] == 1 {
		if k == 0 {
			return -1
		} else {
			k--
		}
	}
	minSteps := math.MaxInt32
	for i := 0; i < len(rows); i++ {
		n := []int{current[0] + rows[i], current[1] + cols[i]}
		if isSafeToGo(grid,visited, n) {
			(*visited)[n[0]][n[1]] = true
			stepsNeeded := shortestPathHelper(grid,visited, k, n, steps+1)
			if stepsNeeded != -1 {
				minSteps = min(minSteps, stepsNeeded)
			}
			(*visited)[n[0]][n[1]] = false
		}
	}
	return minSteps
}

func isSafeToGo(grid [][]int, v *[][]bool, point []int) bool {
	m := point[0]
	n := point[1]
	if m < len(grid) && n < len(grid[m]) && m >= 0 && n >= 0 && !(*v)[m][n]{
		return true
	}
	return false
}
