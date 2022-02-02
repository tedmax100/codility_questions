package main

func maxProfit(prices []int) int {
	var length int = len(prices)
	if length < 1 || length > 30000 {
		return 0
	}

	var profit int = 0

	for i := 1; i < length; i++ {
		if prices[i] > 10000 {
			return 0
		}

		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
