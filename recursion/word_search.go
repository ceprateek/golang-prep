package recursion

import "fmt"

const NumberOfAlphabets = 26

type TrieNode struct {
	children [NumberOfAlphabets]*TrieNode
	terminal bool
}

type Trie struct {
	root *TrieNode
}

func InitTrie() *Trie {
	t := &Trie{root: &TrieNode{}}
	return t
}

func (t *Trie) Insert(word string) bool {
	current := t.root
	for i := 0; i < len(word); i++ {
		charOfWord := word[i]
		indexOfChar := charOfWord - 'a'
		if current.children[indexOfChar] == nil {
			current.children[indexOfChar] = &TrieNode{}
		}
		current = current.children[indexOfChar]
	}
	if current.terminal {
		return false
	}
	current.terminal = true
	return true
}

func (t *Trie) Print(){
	current := t.root
	var result []string
	var temp string
	t.printHelper(&temp, &result, current)
	fmt.Println(result)
}

func (t *Trie) printHelper(temp *string, result *[]string, current *TrieNode){
	if current.terminal{
		r := *temp
		*result = append(*result, r)
	}
	for i:=0;i<len(current.children);i++{
		if current.children[i] != nil{
			currentChar := i + 'a'
			*temp = *temp + string(currentChar)
			t.printHelper(temp, result, current.children[i])
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}

func (t *Trie) SearchPrefix(prefix string) bool {
	current := t.root
	for i := 0; i < len(prefix); i++ {
		charOfWord := prefix[i]
		indexOfChar := charOfWord - 'a'
		if current.children[indexOfChar] == nil {
			return false
		}
		current = current.children[indexOfChar]
	}
	return true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for i := 0; i < len(word); i++ {
		charOfWord := word[i]
		indexOfChar := charOfWord - 'a'
		if current.children[indexOfChar] == nil {
			return false
		}
		current = current.children[indexOfChar]
	}
	if current.terminal {
		return true
	}
	return false
}

var iindex = []int{-1, 0, 0, 1}
var jindex = []int{0, -1, 1, 0}

func isSafeOnBoard(board [][]byte, accessed *[][]bool, i, j int) bool {
	if i >= 0 && j >= 0 && i < len(board) && j < len(board[i]) && !(*accessed)[i][j] {
		return true
	}
	return false
}

func findAllWordsOnBoard(board [][]byte, words []string) []string {
	t := InitTrie()
	for i := 0; i < len(words); i++ {
		t.Insert(words[i])
	}
	result := make([]string, 0, len(words))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			temp := string(board[i][j])
			accessed := make([][]bool, len(board))
			for i := 0; i < len(accessed); i++ {
				accessed[i] = make([]bool, len(board[i]))
			}
			accessed[i][j] = true
			backtrackFromIndex(board, i, j, &temp, &result, t, &accessed)
		}
	}
	return result
}

func backtrackFromIndex(board [][]byte, y, x int, prefix *string, result *[]string, t *Trie, accessed *[][]bool) {
	if t.Search(*prefix) {
		a := *prefix
		*result = append(*result, a)
		return
	} else if t.SearchPrefix(*prefix){
		for i := 0; i < len(iindex); i++ {
			neighborIdxI := y + iindex[i]
			neighborIdxJ := x + jindex[i]
			if isSafeOnBoard(board, accessed, neighborIdxI, neighborIdxJ) {
				(*accessed)[neighborIdxI][neighborIdxJ] = true
				*prefix = *prefix + string(board[neighborIdxI][neighborIdxJ])
				backtrackFromIndex(board, neighborIdxI, neighborIdxJ, prefix, result, t, accessed)
				*prefix = (*prefix)[:len(*prefix)-1]
			}
		}
	}
}

func PlayWordSearch() {
	words := []string{"oath", "pea", "eat", "rain", "peas"}
	t := InitTrie()
	for i := 0; i < len(words); i++ {
		t.Insert(words[i])
	}
	t.Print()
	board := [][]byte{{byte('o'), byte('a'), byte('a'), byte('n')}, {byte('e'), byte('t'), byte('a'), byte('e')}, {byte('i'), byte('h'), byte('k'), byte('r')}, {byte('i'), byte('f'), byte('l'), byte('v')}}
	r := findAllWordsOnBoard(board, words)
	fmt.Println(r)
}
