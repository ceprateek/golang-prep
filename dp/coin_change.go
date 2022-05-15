package dp

import (
	"fmt"
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		for c := 0; c < len(coins); c++ {
			if coins[c] > i {
				continue
			}
			sol := 1 + dp[i-coins[c]]
			dp[i] = min(dp[i], sol)
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PlayCoinChange() {
	coints := []int{1, 2, 5}
	amount := 100
	fmt.Println(coinChange(coints, amount))
}
