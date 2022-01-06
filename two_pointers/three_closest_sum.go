package two_pointers

import (
	"fmt"
	"math"
	"sort"
)

func three_sum_closest(in []int, sum int) {
	closest_sum := math.MaxInt32
	sort.Ints(in)
	for i := 0; i < len(in)-2; i++ {
		left := i + 1
		right := len(in) - 1
		for left < right {
			currentSum := in[i] + in[left] + in[right]
			if math.Abs(float64(currentSum-sum)) < math.Abs(float64(closest_sum-sum)) {
				closest_sum = currentSum
			}
			if currentSum > sum{
				right--
			}else {
				left++
			}
		}
		fmt.Println(closest_sum)
	}
	fmt.Println(closest_sum)
}

func PlayClosestSum()  {
	in := []int{ -1, 2, 1, -4 }
	three_sum_closest(in, 1)
}