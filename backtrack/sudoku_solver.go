package backtrack

import "fmt"

/*
Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],
[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],
["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],
[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],
[".",".",".",".","8",".",".","7","9"]]
Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],
["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],
["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],
["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],
["3","4","5","2","8","6","1","7","9"]]

https://leetcode.com/explore/learn/card/recursion-ii/472/backtracking/2796/

*/

func solve_sudoku(board [][]byte) bool {
	isDone := true
	row := -1
	col := -1
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				isDone = false
				row = i
				col = j
				break
			}
		}
		if !isDone {
			break
		}
	}
	if isDone {
		return isDone
	}
	for i := 1; i < 9; i++ {
		if isSafe(board, row, col, uint8(i)){
			board[row][col] = uint8(i)
			sol := solve_sudoku(board)
			if !sol{
				board[row][col] = 0
			} else {
				fmt.Println(board)
				return true
			}
		}
	}
	return false
}

func isSafe(board [][]byte, i, j int, val uint8) bool {
	if i >= 0 && j >= 0 && i < len(board) && j < len(board[i]) &&
		isSafeInSmallSq(board, i, j, val) && isSafeOnICol(board, i, j, val) && isSafeOnJRow(board, i, j, val) {
		return true
	}
	return false
}

func isSafeOnICol(board [][]byte, i, j int, val uint8) bool {
	for k := 0; i < len(board); k++ {
		if board[k][j] == val {
			return false
		}
	}
	return true
}

func isSafeOnJRow(board [][]byte, i, j int, val uint8) bool {
	for k := 0; i < len(board[i]) && k != j; k++ {
		if board[i][k] == val {
			return false
		}
	}
	return true
}

func isSafeInSmallSq(board [][]byte, i, j int, val uint8) bool {
	starti := (i / 3) * 3
	startj := (j / 3) * 3
	for k := starti; k < starti+3; k++ {
		for l := startj; l < startj+3; l++ {
			if board[k][l] == val {
				return false
			}
		}
	}
	return true
}
