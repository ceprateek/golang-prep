package algorithms

import "fmt"

func PlayMergeSort() {
	arr := []int{12, 3, 18, 24, 0, 5, 2}
	Sort(&arr)
}

func Sort(input *[]int) {
	aux := make([]int, len(*input))
	for i := 0; i < len(*input); i++ {
		aux[i] = (*input)[i]
	}
	mergeSort(input, &aux, 0, len(aux)-1)
	fmt.Println(*input)
}

func mergeSort(input *[]int, aux *[]int, start int, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSort(input, aux, start, mid)
	mergeSort(input, aux, mid+1, end)
	merge(input, aux, start, end, mid)
}

func merge(input *[]int, aux *[]int, start int, end int, mid int) {
	i := start
	j := mid + 1
	k := start
	for i <= mid && j <= end {
		if (*input)[i] < (*input)[j] {
			(*aux)[k] = (*input)[i]
			i++
		} else {
			(*aux)[k] = (*input)[j]
			j++
		}
		k++
	}
	for i <= mid {
		(*aux)[k] = (*input)[i]
		i++
		k++
	}
	for itr:=start; itr<=end; itr++{
		(*input)[itr] = (*aux)[itr]
	}
}
