package arrays

import "fmt"

type out struct {
	num1, num2 int
}

func Play2Sum() {
	packets := []int{1, 5, 7, 2, 8, 4, 3}
	fmt.Println(twoSum(packets, 9))
}
func twoSum(input []int, desiredSum int) (output []out) {
	i := 0
	j := len(input) - 1
	HeapSort(&input)
	for i < j {
		sum := input[i] + input[j]
		if sum == desiredSum {
			output = append(output, out{
				num1: input[i],
				num2: input[j],
			})
			i++
			j--
		} else if sum > desiredSum {
			j--
		} else {
			i++
		}
	}
	return output
}
