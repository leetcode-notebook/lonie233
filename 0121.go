package main

func maxProfit(prices []int) int {

	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	profits := make([]int, len(prices))

	profits[0] = 0
	maxProfit := profits[0]

	minPrice := prices[0]
	for n := 1; n < len(prices); n++ {
		profits[n] = getTwoMaxNumber(profits[n-1], prices[n]-minPrice)
		if profits[n] > maxProfit {
			maxProfit = profits[n]
		}
		minPrice = getTwoMinNumber(minPrice, prices[n])
	}
	return maxProfit
}

func getTwoMaxNumber(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func getTwoMinNumber(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	} else {
		return n1
	}
}

func getMaxProfit(n int, curPrice int, prices []int) int {

	maxProfit := curPrice - prices[0]
	for index := 0; index < n; index++ {
		curProfit := curPrice - prices[index]
		if curProfit > maxProfit {
			maxProfit = curProfit
		}
	}
	return maxProfit
}
