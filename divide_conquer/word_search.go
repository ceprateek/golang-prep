package divide_conquer

import "fmt"

func PlayBoardWordSearch()  {
	//board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	board := [][]byte{{'A','B'},{'C','D'}}
	r := exist(board, "ACDB")
	fmt.Println(r)
}

func exist(board [][]byte, word string) bool {
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[i]);j++ {
			temp := string(board[i][j])
			visited := make([][]bool, len(board))
			for i:=0;i<len(visited);i++{
				visited[i] = make([]bool, len(board[i]))
			}
			visited[i][j] = true
			result := parse(board, word, temp, visited,i,j)
			if result{
				return true
			}
		}
	}
	return false
}

//func returns if the cell can be visited
func canVisit(i,j int, visited [][]bool,board [][]byte) bool{
	if i>=0 && i<len(board) && j>=0 && j<len(board[i]) && !visited[i][j]{
		return true
	}
	return false
}

var row = []int{0,0,-1,1}
var col = []int{-1,1,0,0}

func parse(board [][]byte, word, temp string, visited [][]bool, i,j int) bool{
	if temp == word {
		return true
	}
	lastIdxTemp := len(temp)-1
	if len(temp)>0 && temp[lastIdxTemp] != word[lastIdxTemp]{
		return false
	}
	for idx:=0;idx<len(row);idx++{
		if canVisit(i+row[idx],j+col[idx], visited, board){
			temp1 := temp + string(board[i+row[idx]][j+col[idx]])
			visited[i+row[idx]][j+col[idx]] = true
			result := parse(board, word, temp1, visited,i+row[idx],j+col[idx])
			if result{
				return result
			}
			visited[i+row[idx]][j+col[idx]] = false
		}
	}
	return false
}
