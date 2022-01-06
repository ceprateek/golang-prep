package stacks

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/*
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
*/

/*

the logic is ..
all the digits are in increasing order than we will delete the last digit ,
otherwise we will delete a digit for which next digit is smaller.


"1234567890"
9

*/

func removeKdigits(num string, k int) string {
	result := num
	result = removeDigits(result, k)
	for len(result) > 1 && result[0:1] == "0" {
		runes := []rune(result)
		runes = runes[1:]
		result = string(runes)
	}
	return result
}

func removeDigits(val string, k int) string {
	if len(val) == 1 {
		return "0"
	}
	cache := Stack{}
	var digitFound int
	for i, w := 0, 0; i < len(val); i += w {
		var r rune
		r, w = utf8.DecodeRuneInString(val[i:])
		intStr := string(r)
		val, _ := strconv.Atoi(intStr)
		if cache.IsEmpty() {
			cache.Push(intStr)
			continue
		}
		topOfStack, _ := strconv.Atoi(cache.Peek().(string))
		for val < topOfStack && digitFound < k && !cache.IsEmpty(){
			cache.Pop()
			digitFound++
			if cache.Peek() == nil{
				break
			}
			topOfStack, _ = strconv.Atoi(cache.Peek().(string))
		}
		cache.Push(intStr)
	}
	for digitFound < k && !cache.IsEmpty() {
		cache.Pop()
		digitFound++
	}
	var out string
	for !cache.IsEmpty() {
		out = cache.Pop().(string) + out
	}
	if len(out) == 0 {
		out = "0"
	}
	return out

}

func PlayRemoveKInts() {
	a := "1234567890"
	k := 9
	result := removeKdigits(a, k)
	fmt.Println(result)
}
