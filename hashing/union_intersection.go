package hashing

import (
	"fmt"
	"sort"
)

/*
Given two unsorted arrays that represent two sets
(elements in every array are distinct),
find the union and intersection of two arrays.

For example, if the input arrays are:
arr1[] = {7, 1, 5, 2, 3, 6}
arr2[] = {3, 8, 6, 20, 7}

{1,2,3,5,6,7}
{3,6,7,9,20}

*/

func unionAndIntersection(arr1, arr2 []int) ([]int, []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var intersection, union []int
	var i, j int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			intersection = append(intersection, arr2[j])
			union = append(union, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			if len(union) == 0 || arr1[i] != union[len(union)-1] {
				union = append(union, arr1[i])
			}
			i++
		} else {
			if len(union) == 0 || arr2[i] != union[len(union)-1] {
				union = append(union, arr2[i])
			}
			j++
		}
	}
	for i < len(arr1) {
		union = append(union, arr1[i])
		i++
	}
	for j < len(arr2) {
		union = append(union, arr2[j])
		j++
	}

	return union, intersection
}

func PlayUnionIntersection() {
	arr1 := []int{7, 1, 5, 2, 3, 6}
	arr2 := []int{3, 8, 6, 20, 7}

	union, intersection := unionAndIntersection(arr1, arr2)
	fmt.Println(union)
	fmt.Println(intersection)
}
