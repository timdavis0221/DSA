package dsa

import "fmt"

// MaxProfit returns a max profit from a given array.
func MaxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit < 0 {
				continue
			}
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// MaxProfitSolution2 makes the O(n) improvements to return the maximum profit from the given array.
func MaxProfitSolution2(prices []int) int {
	maxProfit, currentProfit := 0, 0
	buy := prices[0]

	for _, sell := range prices[1:] {
		if sell > buy {
			currentProfit = sell - buy
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		} else {
			// Sell price < buy price, so update buy price.
			// Because the lower buy price always give the higher profits.
			buy = sell
		}
	}
	fmt.Println("Max profit: ", maxProfit)
	return maxProfit
}
