package stacks

import (
	"fmt"
	"unicode/utf8"
)

/*
Input: exp = “[()]{}{[()()]()}”
Output: Balanced

Input: exp = “[(])”
Output: Not Balanced
*/

func PlayParenthesisBalance() {
	exp := "[()]{}{[()()]()}"
	fmt.Println(checkBalance(exp))
}

func checkBalance(input string) bool {
	stack := Stack{}
	for i, w := 0, 0; i < len(input); i += w {
		r, width := utf8.DecodeRuneInString(input[i:])
		w = width
		if r == '[' || r == '(' || r == '{' {
			stack.Push(r)
			continue
		}
		if stack.IsEmpty() {
			return false
		}
		switch r {
		case ')':
			val := stack.Pop()
			if val != '(' {
				return false
			}
		case '}':
			val := stack.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := stack.Pop()
			if val != '[' {
				return false
			}
		}
	}
	return true
}
