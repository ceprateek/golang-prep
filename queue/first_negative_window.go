package queue

import "fmt"

/*
Input : arr[] = {-8, 2, 3, -6, 10}, k = 2
Output : -8 0 -6 -6
First negative integer for each window of size k
{-8, 2} = -8
{2, 3} = 0 (does not contain a negative integer)
{3, -6} = -6
{-6, 10} = -6

Input : arr[] = {12, -1, -7, 8, -15, 30, 16, 28} , k = 3
Output : -1 -1 -7 -15 -15 0
*/

func firstNegativeNumberInWindow(arr []int, window int) []int {
	var out []int
	q := Queue{}
	//12 -1 -7
	for i := 0; i < len(arr); i++ {
		q.Add(arr[i])
		if q.Size() < window {
			continue
		}
		itr := q.head
		var found bool
		for itr != nil {
			data := itr.data
			if data.(int) < 0 {
				out = append(out, data.(int))
				found = true
				break
			}
			itr = itr.next
		}
		if !found{
			out = append(out, 0)
		}
		q.Remove()
	}
	return out
}

func PlayFirstNegNumWin() {
	arr := []int{12, -1, -7, 8, -15, 30, 16, 28}
	out := firstNegativeNumberInWindow(arr, 3)
	fmt.Println(out)
}
