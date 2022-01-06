package sorting

import (
	"fmt"
	"sort"
)

func PlayCountSort()  {
	in := []int{0,0,1,2,1}
	fmt.Println(count_sort(in, 2))
	sort.Ints(in)
	fmt.Println(in)
}

func count_sort(in []int, max int) []int{
	count := make([]int, max+1)
	for _,val := range in{
		count[val]++
	}
	fmt.Println(count)
	for i:=1; i<len(count); i++{
		count[i] += count[i-1]
	}
	fmt.Println(count)
	out := make([]int, len(in))

	for i:=0; i<len(in); i++{
		idx := count[in[i]]
		fmt.Println(idx)
		out[idx-1] = in[i]
		fmt.Println(out)
		count[in[i]]--
		fmt.Println(count)
	}
	return out
}
