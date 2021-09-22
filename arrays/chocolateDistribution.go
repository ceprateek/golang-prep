package arrays

import (
	"fmt"
	"math"
)

func PlayChocolateDistribution() {
	packets := []int{5, 7, 2, 8, 4, 3}
	distributeChocolates(&packets, 3)
}

func distributeChocolates(packets *[]int, students int) {
	HeapSort(packets)
	val := *packets
	fmt.Println(val)
	minDiff := math.MaxInt32
	for i := 0; i <= len(val)-students; i++ {
		diff := val[i+students-1] - val[i]
		fmt.Println(fmt.Sprintf("diff between: %d %d", i, i+students-1))
		if diff < minDiff {
			minDiff = diff
		}
	}
	fmt.Println(minDiff)
}

func HeapSort(in *[]int) {
	Heapify(in)
	val := *in
	lastIdx := len(val) - 1
	for i := 0; i < len(val); i++ {
		val[0], val[lastIdx] = val[lastIdx], val[0]
		lastIdx--
		HeapifyAtIdx(in, lastIdx, 0)
	}
}

func Heapify(in *[]int) {
	for i := len(*in) / 2; i >= 0; i-- {
		HeapifyAtIdx(in, len(*in)-1, i)
	}
}

func HeapifyAtIdx(in *[]int, lastIndex int, idx int) {
	arr := *in
	left := 2*idx + 1
	right := 2*idx + 2
	max := idx
	if left <= lastIndex && arr[left] > arr[idx] {
		max = left
	}
	if right <= lastIndex && arr[right] > arr[max] {
		max = right
	}
	if max != idx {
		arr[idx], arr[max] = arr[max], arr[idx]
		HeapifyAtIdx(in, lastIndex, max)
	}
}
