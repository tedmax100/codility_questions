package main

func maxProfit(prices []int) int {
	var length int = len(prices)
	if length < 1 || length > 100000 {
		return 0
	}

	var profit int = 0
	var buyPrice int = 100000

	for i := 0; i < len(prices); i++ {
		if prices[i] > 10000 {
			return 0
		}
		buyPrice = min(buyPrice, prices[i])
		profit = max(profit, prices[i]-buyPrice)
	}

	return profit
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
