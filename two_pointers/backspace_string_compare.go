package two_pointers

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func PlayBackspaceCompare() {
	s1 := "geee#e#ks"
	s2 := "gee##eeks"
	result := backspaceCompare(s1, s2)
	fmt.Println(result)
}

func backspaceCompare(s1, s2 string) bool {
	s1Clean := removeBackspace(s1)
	s2Clean := removeBackspace(s2)
	fmt.Printf("%s %s\n", s1Clean, s2Clean)
	if strings.Compare(s1Clean, s2Clean) == 0{
		return true
	}
	return false
}

func removeBackspace(in string) string{
	var s stack
	for i,w := 0,0; i<len(in); i+=w {
		rune, width := utf8.DecodeRuneInString(in[i:])
		if rune == '#'{
			s.Pop()
		}else {
			s.Push(string(rune))
		}
		w = width
	}
	return s.getString()
}

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() {
	index := len(*s) - 1
	*s = (*s)[:index]
}

func (sptr *stack) getString() string {
	var out string
	for _,str := range *sptr{
		out = out +  str
	}
	return out
}
