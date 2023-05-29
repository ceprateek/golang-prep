package backtrack

import "fmt"

func PrintDiceSum(length, sum int) {
	printDiceSumHelper(length, sum, 0, []int{})
}

func printDiceSumHelper(length, sum, tempSum int, temp []int) {
	if length == 0 {
		if sum == tempSum {
			fmt.Println(temp)
		}
	} else {
		for i := 1; i <= 6; i++ {
			tempSum += i
			temp = append(temp, i)
			printDiceSumHelper(length-1, sum, tempSum, temp)
			tempSum -= i
			temp = temp[:len(temp)-1]
		}
	}
}
