package dp

import "fmt"

func PlayUglyNumber(){
	n := nthUglyNumber(11)
	fmt.Println(n)
}

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m2:=1
	m3:=1
	m5:=1

	for i:=2;i<=n;i++{
		v2 := 2*dp[m2]
		v3 := 3*dp[m3]
		v5 := 5*dp[m5]
		dp[i] = min(v2, min(v3, v5))
		if dp[i] == v2{
			m2++
		}
		if dp[i] == v3 {
			m3++
		}
		if dp[i] == v5 {
			m5++
		}
		fmt.Printf("i:%d dp[i]:%d v2:%d v3:%d v5:%d m2:%d m3:%d m5:%d\n", i, dp[i], v2, v3, v5, m2, m3, m5)
	}

	return dp[n]

}

