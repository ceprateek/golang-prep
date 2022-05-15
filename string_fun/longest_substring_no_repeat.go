package string_fun

import "fmt"

func PlayLongestSubstring(){
	in := "geeksforgeeks"
	r := findLongestWithoutRepeatingChars(in)
	fmt.Println(r)
}

func findLongestWithoutRepeatingChars(in string) string {
	var result string
	for i:=0;i<len(in);i++{
		for j:=i+1;j<len(in);j++{
			sub := in[i:j+1]
			if checkIfStringContainsUniqueChars(sub) && len(sub)>len(result){
				result = sub
			}
		}
	}
	return result
}

func checkIfStringContainsUniqueChars(in string) bool{
	if len(in) == 0{
		return true
	}
	contains := make([]bool, 26)
	for i := 0; i < len(in); i++ {
		idx := in[i]-'a'
		if contains[idx]{
			return false
		}
		contains[idx] = true
	}
	return true
}
