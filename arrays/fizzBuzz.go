package arrays

import (
	"fmt"
	"strconv"
)

/*
Write a short program that prints each number from 1 to 100 on a new line.
For each multiple of 3, print “Fizz” instead of the number.
For each multiple of 5, print “Buzz” instead of the number.
For numbers that are multiples of both 3 and 5, print “FizzBuzz” instead of the number.
*/

func PlayPrintFizzBuzz() {
	printFizzBuzz(1, 100)
}

func printFizzBuzz(begin, end int) {
	for i := begin; i < end; i++ {
		out := ""
		if i%3 == 0 {
			out = "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if len(out) == 0 {
			out = strconv.Itoa(i)
		}
		fmt.Println(out)
	}
}
