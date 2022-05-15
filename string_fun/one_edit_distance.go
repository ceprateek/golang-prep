package string_fun

import (
	"fmt"
)

func PlayMidEditDistance() {
	s := "horse"
	t := "ros"
	minDist := minEditDistance(s, t)
	fmt.Println(minDist)
}
func minEditDistance(s, t string) int {
	return minEditDistanceHelper(s, t, len(s)-1, len(t)-1, 0)
}


func minEditDistanceHelper(s, t string, sidx, tidx, currentEdits int) int {
	if sidx < 0 && tidx < 0 {
		return currentEdits
	}
	if sidx < 0 {
		return currentEdits + tidx + 1
	}
	if tidx < 0 {
		return currentEdits + sidx + 1
	}

	if s[sidx] == t[tidx] {
		return minEditDistanceHelper(s, t, sidx-1, tidx-1, currentEdits)
	}

	//1. Insert in s
	insertDist := minEditDistanceHelper(s, t, sidx, tidx-1, currentEdits+1)
	//2. remove in s
	delDist := minEditDistanceHelper(s, t, sidx-1, tidx, currentEdits+1)
	//3. replace in s
	replDist := minEditDistanceHelper(s, t, sidx-1, tidx-1, currentEdits+1)

	return min(insertDist, delDist, replDist)
}

func min(a, b, c int) int {
	var r int
	if a < b {
		r = a
	} else {
		r = b
	}
	if r < c {
		return r
	} else {
		return c
	}
}

//abc abd
