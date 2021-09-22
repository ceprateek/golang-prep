package arrays

import "fmt"

func PrintSubstrings(in string) {
	N := len(in)
	for length := 1; length <= N; length++ {
		for start := 0; start <= N-length; start++ {
			fmt.Println(in[start:start+length])
		}
	}
}
