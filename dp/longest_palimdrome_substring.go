package dp

import "fmt"

func PlayLongestPalindrone() {
	s := "cbbd"
	r := longestPalindrome(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		l1 := expandFromCenter(s, i, i)
		l2 := expandFromCenter(s, i, i+1)
		l := max(l1, l2)
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandFromCenter(s string, start, end int) int {
	left := start
	right := end

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
