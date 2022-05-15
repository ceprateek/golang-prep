package recursion

import "fmt"

/*
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]

*/

func PlayGenParenthesis()  {
	genParenthesis(3)
}

func genParenthesis(n int) {
	result := make([]string,n)
	temp := "()"
	genParenthesisHelper(n-1, &temp, &result)
	fmt.Println(result)
}

func genParenthesisHelper(n int, temp *string, result *[]string) {
	if n == 0 {
		if !sliceContainsString(*result, *temp) {
			*result = append(*result, *temp)
		}
	} else {
		for i := 0; i < len(*temp); i++ {
			originalTemp := *temp
			*temp = (*temp)[:i] + "()" + (*temp)[i:]
			genParenthesisHelper(n-1, temp, result)
			temp = &originalTemp
		}
	}
}

func sliceContainsString(in []string, s string) bool {
	for _, val := range in {
		if val == s {
			return true
		}
	}
	return false
}
