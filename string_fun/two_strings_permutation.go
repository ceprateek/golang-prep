package string_fun

import "fmt"

func checkTwoStrintPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	cache := make([]int, 256)
	for i := 0; i < len(s1); i++ {
		cache[s1[i]]++
	}

	for i := 0; i < len(s2); i++ {
		cache[s2[i]]--
		if cache[s2[i]] < 0{
			return false
		}
	}

	return true
}

func PlayCheckTwoStringsPermutate() {
	s := "abcaab"
	b := "bcaaac"
	result := checkTwoStrintPermutation(s, b)
	fmt.Println(result)

}
