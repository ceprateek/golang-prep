package hashing

import (
	"fmt"
	"sort"
)

/*
Given an array of n numbers and a positive integer k.
The problem is to find k numbers with most occurrences, i.e.,
the top k numbers having the maximum frequency.
If two numbers have the same frequency then the larger number
should be given preference.
The numbers should be displayed in decreasing order of their frequencies.
It is assumed that the array consists of k numbers with most occurrences.

Input:
arr[] = {3, 1, 4, 4, 5, 2, 6, 1},
k = 2
Output: 4 1
Explanation:
Frequency of 4 = 2
Frequency of 1 = 2
These two have the maximum frequency and
4 is larger than 1.

Input :
arr[] = {7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9},
k = 4
Output: 5 11 7 10
Explanation:
Frequency of 5 = 3
Frequency of 11 = 2
Frequency of 7 = 2
Frequency of 10 = 1
These four have the maximum frequency and
5 is largest among rest.
*/

func numbersKFrequency(input []int, k int) {
	cache := make(map[int]int)
	keys := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if _, ok := cache[input[i]]; ok {
			cache[input[i]] += 1
		} else {
			cache[input[i]] = 1
			keys = append(keys, input[i])
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		if cache[keys[i]] > cache[keys[j]] {
			return true
		} else if cache[keys[i]] == cache[keys[j]] {
			return keys[i] > keys[j]
		} else {
			return false
		}
	})
	for i := 0; i < k; i++ {
		fmt.Printf("%d has frequency: %d\n", keys[i], cache[keys[i]])
	}
}

func PlayKFrequency() {
	in := []int{7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9}
	k := 4
	numbersKFrequency(in, k)
}
