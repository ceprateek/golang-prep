package backtrack

import (
	"fmt"
	"testing"
)

func solveSudoku(board [][]byte) bool {
	return solveSudokuUtil(board, 0, 0)
}

func solveSudokuUtil(board [][]byte, row, col int) bool {
	if row == len(board) {
		return true
	}

	if col == len(board[0]) {
		return solveSudokuUtil(board, row+1, 0)
	}

	if board[row][col] != '.' {
		return solveSudokuUtil(board, row, col+1)
	}

	for num := byte('1'); num <= '9'; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num
			if solveSudokuUtil(board, row, col+1) {
				return true
			}
			board[row][col] = '.'
		}
	}

	return false
}

func isSafe(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

func TestSudokuSolver(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	expected := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != expected[i][j] {
				t.Errorf("Expected %c at [%d][%d], but got %c", expected[i][j], i, j, board[i][j])
			}
		}
	}
	fmt.Println(board)
}

func RunTestSudokuSolver() {
	tests := []testing.InternalTest{
		{Name: "TestSudokuSolver", F: TestSudokuSolver},
	}
	testing.Main(nil, tests, nil, nil)
}
