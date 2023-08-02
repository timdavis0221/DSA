package dsa

// MaxProfit returns a max profit from a give array.
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
