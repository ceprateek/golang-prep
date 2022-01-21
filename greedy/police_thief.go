package greedy

import "fmt"

/*
https://www.geeksforgeeks.org/policemen-catch-thieves/

Input : arr[] = {'P', 'T', 'T', 'P', 'T'},
            k = 1.
Output : 2.
Input : arr[] = {'T', 'T', 'P', 'P', 'T', 'P'},
            k = 2.
Output : 3.

Input : arr[] = {'P', 'T', 'P', 'T', 'T', 'P'},
            k = 3.
Output : 3.
*/

/*
sol:
1. Get the lowest index of policeman p and thief t. Make an allotment
if |p-t| <= k and increment to the next p and t found.
2. Otherwise increment min(p, t) to the next p or t found.
3. Repeat above two steps until next p and t are found.
4. Return the number of allotments made.
*/

func PlayPoliceThief() {
	in := []string{"P", "T", "P", "T", "T", "P"}
	k := 3
	fmt.Println(policeCatchThief(in, k))
}

func policeCatchThief(in []string, k int) int {
	result := 0
	police := make([]int, 0)
	thief := make([]int, 0)
	for i := 0; i < len(in); i++ {
		if in[i] == "P" {
			police = append(police, i)
		} else {
			thief = append(thief, i)
		}
	}
	var p, t int
	for p < len(police) && t < len(thief) {
		if Abs(thief[t]-police[p])<=k{
			t++
			p++
			result++
		}else if thief[t]<police[p]{
			t++
		}else {
			p++
		}
	}
	return result
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
