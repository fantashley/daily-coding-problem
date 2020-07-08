package main

import "fmt"

func main() {
	stockPrices := []int{10, 7, 5, 8, 11, 9}
	fmt.Println(getMaxProfit(stockPrices))
}

func getMaxProfit(stockPrices []int) int {
	buy := stockPrices[0]
	maxProfit := stockPrices[1] - buy
	for _, price := range stockPrices[1:] {
		if price < buy {
			buy = price
		} else if (price - buy) > maxProfit {
			maxProfit = price - buy
		}
	}
	return maxProfit
}
