package arrays

import (
	"fmt"
	"math"
)

func PlayStocksSingle() {
	input := []int{3, 5, 1, 7, 4, 9, 3}
	fmt.Println(returnMaxProfitUnlimitedTransactions(input))
}
func returnMaxProfitUnlimitedTransactions(stockPrices []int) int{
	if len(stockPrices) <= 1 {
		return 0
	}
	i := 0
	l := len(stockPrices) - 1
	var maxProfit int
	for i < l {
		for i<l && stockPrices[i+1] < stockPrices[i] {
			i++
		}
		if i==l{
			break
		}
		localmin := stockPrices[i]
		for i<l && stockPrices[i+1] > stockPrices[i] {
			i++
		}
		localmax := stockPrices[i]
		maxProfit += localmax-localmin

	}
	return maxProfit
}

func returnMaxProfit(stockPrices []int) int {
	maxProfit := 0
	currentProfit := 0
	minPriceTillNow := math.MaxInt32
	for i := 0; i < len(stockPrices); i++ {
		if minPriceTillNow > stockPrices[i] {
			minPriceTillNow = stockPrices[i]
		}
		currentProfit = stockPrices[i] - minPriceTillNow
		if maxProfit < currentProfit {
			maxProfit = currentProfit
		}
	}
	return maxProfit
}
