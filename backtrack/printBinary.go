package backtrack

import "fmt"

func PrintBinary() {
	printBinaryHelper(3, "")
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
