package searching

import (
	"fmt"
	"sort"
)

/*
Input: n = 5, k = 4
nums[] = {2, 3, 6, 2, 4}
Output: 3
Explanation: Wood collected by cutting trees at height 3 = 0 + 0 + (6-3) + 0 + (4-3) = 4
hence 3 is to be subtracted from all numbers. Since 2 is lesser than 3, nothing gets
subtracted from it.
*/

/*
sort
2,2,3,4,6

wood: 4

*/

func Play_height() {
	in := []int{4, 42, 40, 26, 46}
	h := find_height(in, 20)
	fmt.Println(h)
}

func find_height(in []int, K int) (height int) {
	sort.Ints(in)
	height = -1
	l := 0
	r := in[len(in)-1]
	for l < r {
		mid := l + (r-l)/2
		woodsAtMid := cutAtHeight(in, mid)
		if woodsAtMid == K {
			return mid
		} else if woodsAtMid > K {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return height
}

func cutAtHeight(input []int, height int) (woods int) {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] <= height {
			break
		}
		woods += input[i] - height
	}
	return woods
}
