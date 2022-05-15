package dp

import "fmt"

func PlayWordBreak() {
	dictionary := []string{"cats", "dog", "sand", "and", "cat"}
	s := "catsandog"
	fmt.Println(wordBreak(s, dictionary))
}

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		dict[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if val, ok := dict[s[j:i]]; ok && val && dp[j]{
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
