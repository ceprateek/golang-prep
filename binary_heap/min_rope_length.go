package binary_heap

import "fmt"

/*
(2, 3, 4, 6) the sum is equal to 2 * 3 + 3 * 3 + 4 * 2 + 6 * 1 = 29 (According to the diagram).
*/

func PlayMinRope(){
	out := minRopeLength([]int{2,3,4,6})
	fmt.Println(out)
}

func minRopeLength(in []int) int {
	h := &Heap{}
	for _, val := range in {
		*h = append(*h, val)
	}
	h.Heapify()
	var cost int
	for len(*h) > 1 {
		one := h.Del()
		two := h.Del()
		combination := one + two
		h.Insert(combination)
		cost += combination
		fmt.Println(cost)
	}
	return cost
}
