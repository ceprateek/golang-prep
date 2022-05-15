package algorithms

import "fmt"

func PlayKMP() {
	KMPSearch("AAACAAAA", "ABCABAABCABAC")
}

func KMPSearch(pat string, txt string) {
	M := len(pat)
	N := len(txt)

	lps := make([]int, M)
	i := 0
	j := 0

	computeLps(pat, M, &lps)
	for i < N {
		if pat[j] == txt[i] {
			i++
			j++
		}
		if j == M {
			fmt.Println(fmt.Sprintf("patten at: %d", i-j))
			j = lps[j-1]
		} else if i < N && pat[j] != txt[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
}

func computeLps(pattern string, M int, lps *[]int) {
	//length of previous longest prefix
	length := 0
	i := 1
	(*lps)[0] = 0

	for i < M {
		if pattern[i] == pattern[length] {
			length++
			(*lps)[i] = length
			i++
		} else {
			// This is tricky. Consider the example.
			// AAACAAAA and i = 7. The idea is similar
			// to search step.
			if length != 0 {
				length = (*lps)[length-1]
				// Also, note that we do not increment
				// i here
			} else {
				(*lps)[i] = length
				i++
			}
		}
	}
}
