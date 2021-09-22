package searching

import "fmt"

func PlayMedian2SortedArrays() {
	ar1 := []int{2, 3, 5, 8}
	ar2 := []int{10, 12, 14, 16, 18, 20}
	if len(ar1) == len(ar2) {
		fmt.Printf("same length median1: %d\n", findMedianSameLengthEfficient(ar1, ar2))
		fmt.Printf("same length median2: %d\n", findMedianSameLength(ar1, ar2))
	} else {
		fmt.Printf("diff length median: %d", findMedianDifferentLength(ar1, ar2))
	}

}

func findMedianSameLength(in1, in2 []int) int {
	var m1, m2, i, j, counter int

	for counter <= len(in1) {

		if in1[i] < in2[j] {
			m1, m2 = in1[i], m1
			i++
		} else {
			m1, m2 = in2[j], m1
			j++
		}
		counter++
	}

	return (m1 + m2) / 2
}

func findMedianSameLengthEfficient(in1, in2 []int) int {
	return findMedianSameLengthEfficientHelper(in1, in2, 0, len(in1)-1, 0, len(in2)-1)
}

func findMedianSameLengthEfficientHelper(in1, in2 []int, start1, end1, start2, end2 int) int {
	fmt.Printf("start:%d end:%d start:%d end:%d\n", start1, end1, start2, end2)
	if end1-start1 == 1 {
		fmt.Printf("finale start:%d end:%d start%d end:%d\n", start1, end1, start2, end2)
		fmt.Println(max(in1[start1], in2[start2]))
		fmt.Println(min(in1[end1], in2[end2]))
		return (max(in1[start1], in2[start2]) + min(in1[end1], in2[end2])) / 2
	}
	m1 := getMedian(in1, start1, end1)
	m2 := getMedian(in2, start2, end2)

	if m1 == m2 {
		return m1
	}
	if m1 < m2 {
		return findMedianSameLengthEfficientHelper(in1, in2, (start1+end1+1)/2, end1, start2, (end2+start2)/2)
	} else {
		return findMedianSameLengthEfficientHelper(in1, in2, start1, (start1+end1)/2, (end2+start2+1)/2, end2)
	}
}

func getMedian(arr []int, start, end int) int {
	size := end - start + 1
	if size%2 != 0 {
		return arr[start+size/2]
	} else {
		return (arr[start+size/2] + arr[start+size/2-1]) / 2
	}
}

func findMedianDifferentLength(in1, in2 []int) int {
	return findMedianDiffLengthEfficientHelper(in1, in2, 0, len(in1)-1, 0, len(in2)-1)
}

//TODO implement me
func findMedianDiffLengthEfficientHelper(in1, in2 []int, start1, end1, start2, end2 int) int {
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
