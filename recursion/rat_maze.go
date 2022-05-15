package recursion

import "fmt"

var xaxis = [...]int{-1, 0, 0, 1}
var yaxis = [...]int{0, -1, 1, 0}

func PlayRatMaze(){
	maze := [][]int{ { 1, 0, 0, 0 },
		{ 1, 1, 0, 1 },
		{ 0, 1, 0, 0 },
		{ 1, 1, 1, 1 } }
	solvemaze(maze)
}

func solvemaze(maze [][]int) {
	sol := make([][]int,len(maze))
	for i:=0;i<len(maze);i++{
		sol[i] = make([]int, len(maze[i]))
	}

	if !solveMazeHelper(maze, &sol, 0, 0) {
		fmt.Println(-1)
	} else {
		fmt.Println(sol)
	}
}

func isSafe(maze [][]int, x, y int) bool {
	if x < len(maze[0]) && y < len(maze) && x >= 0 && y >= 0 && maze[y][x] == 1 {
		return true
	}
	return false
}

func solveMazeHelper(maze [][]int, sol *[][]int, x, y int) bool {
	if x == len((*sol)[0]) && y == len(*sol) {
		(*sol)[y][x] = 1
		return true
	}
	if isSafe(maze, x, y) {
		if (*sol)[y][x] == 1 {
			return false
		}
		(*sol)[y][x] = 1
		for i := 0; i < len(xaxis); i++ {

			result := solveMazeHelper(maze, sol, x+xaxis[i], y+yaxis[i])
			if result {
				return true
			}
		}
		(*sol)[y][x] = 0
	}
	return false
}
