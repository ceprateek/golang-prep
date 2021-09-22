package arrays

import "fmt"

func PlaySort012() {
	in := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	sortArray012(&in)
	fmt.Println(in)
}

func sortArray012(a *[]int) {
	b := *a
	lo := 0
	hi := len(*a) - 1
	mid := 0
	for i := 0; i < len(*a) && mid<=hi; i++ {
		fmt.Println(fmt.Sprintf("i: %d lo: %d mid: %d hi: %d a: %v", i, lo, mid, hi, b))
		for b[hi] == 2 {
			hi--
			fmt.Println(fmt.Sprintf("pre 2 i: %d lo: %d mid: %d hi: %d a: %v", i, lo, mid, hi, b))
		}
		if b[i] == 2 {
			b[hi], b[i] = b[i], b[hi]
			hi--
			fmt.Println(fmt.Sprintf("2 i: %d lo: %d mid: %d hi: %d a: %v", i, lo, mid, hi, b))
		}
		if b[i] == 1 {
			b[mid], b[i] = b[i], b[mid]
			mid++
			fmt.Println(fmt.Sprintf("1 i: %d lo: %d mid: %d hi: %d a: %v", i, lo, mid, hi, b))
		}
		if b[i] == 0 {
			b[lo], b[i] = b[i], b[lo]
			lo++
			mid++
			fmt.Println(fmt.Sprintf("0 i: %d lo: %d mid: %d hi: %d a: %v", i, lo, mid, hi, b))
		}
	}
}
