package backtrack

import "fmt"

func PrintBinary(len int) {
	printBinaryHelper(len, "")
}

func printBinaryHelper(length int, temp string) {
	if length == 0 {
		fmt.Println(temp)
	} else {
		temp = temp + "0"
		printBinaryHelper(length-1, temp)
		temp = temp[:len(temp)-1]
		temp = temp + "1"
		printBinaryHelper(length-1, temp)
		temp = temp[:len(temp)-1]
	}

}
