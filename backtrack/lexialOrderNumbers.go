package backtrack

import "fmt"

func PlayLexialOrder() {
	n := 13
	result := lexialOrder(n)
	fmt.Println(result)
}

func lexialOrder(n int) []int {
	result := make([]int, 0)
	lexialOrderHelper(&result, 0, n)
	return result
}

//n=13
//1, 10, 11, 12, 2, 3, 4, 5, 6, 7, 8, 9

//n=21
//1,10---19, 2, 20, 21, 3, ---

func lexialOrderHelper(results *[]int, start int, maxNumber int) {
	if start > maxNumber {
		return
	}

	*results = append(*results, start)

	for i := 0; i < 10; i++ {
		if start == 0 && i == 0 {
			continue
		}
		if start*10+i < maxNumber {
			lexialOrderHelper(results, start*10+i, maxNumber)
		} else {
			break
		}
	}
}
