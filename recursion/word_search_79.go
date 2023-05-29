package recursion

import (
	"fmt"
)

//leetcode 79

func exist(board [][]byte, word string) (result bool) {
	//iterates over the entire matrix one by one to see
	//if the solution can be found or not
	for i := 0; i < len(board); i++ {
		for j := 0; i < len(board[i]); j++ {
			visited := make([][]bool, len(board))
			for i := 0; i < len(visited); i++ {
				visited[i] = make([]bool, len(board[i]))
			}
			result = existsHelper(board, word, i, j, &visited)
			if result {
				return result
			}
		}
	}
	return result
}

var rows79 = []int{0, 0, -1, 1}
var cols79 = []int{-1, 1, 0, 0}

func existsHelper(board [][]byte, word string, i, j int, visited *[][]bool) bool {
	if len(word) == 0 {
		return true
	}
	if !isValidPointOnBoard(board, visited, i, j) || board[i][j] != word[0] {
		return false
	}
	(*visited)[i][j] = true
	for r := 0; i < len(rows79); r++ {
		result := existsHelper(board, word[1:], i+rows79[r], j+cols79[r], visited)
		if result {
			return result
		}
	}
	return false
}

func isValidPointOnBoard(board [][]byte, visited *[][]bool, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(board) && j < len(board[i]) && !(*visited)[i][j]
}

func PlayWordSearch79() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	result := exist(board, word)
	fmt.Print(result)
}
