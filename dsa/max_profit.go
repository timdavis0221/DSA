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

// MaxProfitSolution2 makes the O(n) improvements by Greedy Algorithm
// to return the maximum profit from the given array.
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

// MaxProfitSolution3 finds the maximun profit by Dynamic Programming (Tabulation approach)
// from the given prices array.
func MaxProfitSolution3(prices []int) int {
	length := len(prices)
	// dp[i][0]: the i-th day to hold the stock.
	// dp[i][1]: the i-th day to drop the stock.
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	// The formula of dp[i][0] = max(dp[i-1][0], your original money - prices[i]).
	// dp[i-1][0] means you already bought the stock before i-th day, so just keep it.
	// Also, we could buy the stock at i-th day, so assume the original money is $0
	// or amount of money (dp[0][0]), another possible profit is 0 minus the i-th day price.

	// The formula of dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0]).
	// dp[i-1][1] assumes you alraedy sold the stock before i-th day.
	// You could sell the stock on the i-th day as well, so the profit can be
	// prices[i] plus your original money (dp[i-1][0]).

	// Initilization of dp[0][0] and dp[0][1].
	// dp[0][0]: the i-th day that you buy (hold) the stock.
	// dp[0][1]: the i-th day that you sell (don't hold) the stock.
	dp[0][0] = -prices[0] // dp[0][0] - prices[0]
	dp[0][1] = 0

	// Iterate the table.
	// TODO: space complexity can be optimized.
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}
	return dp[length-1][1]
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
