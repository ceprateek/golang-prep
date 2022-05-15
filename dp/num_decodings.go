package dp

import (
	"fmt"
	"strconv"
)

func PlayNumDecodings() {
	a := numDecodings("126")
	fmt.Println(a)
}

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	firstDigit, _ := strconv.Atoi(s[:1])
	if firstDigit != 0{
		dp[1]=1
	}
	for i:=2;i<=len(s);i++{
		lastDigit, _ := strconv.Atoi(s[i-1:i])
		if lastDigit !=0{
			dp[i] = dp[i-1]
		}

		lastTwoDigits, _ := strconv.Atoi(s[i-2:i])
		if lastTwoDigits>=10 && lastTwoDigits<=26{
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}