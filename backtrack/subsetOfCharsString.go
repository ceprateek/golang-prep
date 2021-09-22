package backtrack

import "fmt"

func PrintSubstring(s string){
	substringHelper(s, "")
}

func substringHelper(s, temp string){
	if len(s) == 0{
		fmt.Println(temp)
	}else {
		chosen := s[0]
		s = s[1:]
		substringHelper(s, temp)
		temp = temp + string(chosen)
		substringHelper(s, temp)
		temp = temp[:len(temp)-1]
		s = string(chosen) + s
	}
}