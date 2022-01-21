package hashing

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

/*
Input: array = {10, 2, 3, 4, 5, 9, 7, 8}
       X = 23
Output: 3 5 7 8
Sum of output is equal to 23,
i.e. 3 + 5 + 7 + 8 = 23.

Input: array = {1, 2, 3, 4, 5, 9, 7, 8}
       X = 16
Output: 1 3 5 7
Sum of output is equal to 16,
i.e. 1 + 3 + 5 + 7 = 16.
*/

func fourSum(inputs []int, sum int) [][]int {
	output := make(map[int64][]int)
	cache := make(map[int][]int)
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs); j++ {
			in := make([]int, 2)
			in[0] = i
			in[1] = j
			cache[inputs[i]+inputs[j]] = in
		}
	}
	for k := range cache {
		if val, ok := cache[sum-k]; ok {
			out := make([]int, 0, 4)
			out = append(out, val...)
			out = append(out, cache[k]...)
			hashOut := getHash(out)
			if _, ok := output[hashOut]; !ok {
				output[hashOut] = out
			}
		}
	}

	var o [][]int
	for _, val := range output {
		o = append(o, val)
	}
	return o
}

func Play4Sum() {
	in := []int{1, 2, 3, 4, 5, 9, 7, 8}
	out := fourSum(in, 16)
	fmt.Println(out)
}

func getHash(in []int) int64 {
	var thash int64
	for _, val := range in {
		h := hash(strconv.Itoa(val))
		thash += int64(h)
	}
	return thash
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
