package recursion

func maxProfit(prices []int) int {
 return maxProfitHelper(prices, false, 0, 0, 0)
}

func maxProfitHelper(prices []int, isbought bool, currentDay, priceBought, profit int) int{
	if currentDay >= len(prices){
		return profit
	}

	if isbought{ //sell block
		//sell
		var p1 int
		currentProfit := prices[currentDay]-priceBought
		if currentProfit>0{
			p1 =  maxProfitHelper(prices, false, currentDay+1, 0, profit+currentProfit)
		}

		//dont sell
		p2 := maxProfitHelper(prices, true, currentDay+1, priceBought, profit)
		return max(p1, p2)
	} else { //buy block
		//buy
		p1 := maxProfitHelper(prices,true, currentDay+1, prices[currentDay], profit)
		//dony buy
		p2 := maxProfitHelper(prices,false, currentDay+1, 0, profit)
		return max(p1, p2)

	}

}

func max(a, b int) int{
	if a>b {
		return a
	}
	return b
}
