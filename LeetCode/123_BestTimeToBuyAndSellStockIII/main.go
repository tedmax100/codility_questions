package main

func maxProfit(prices []int) int {
	var length int = len(prices)
	if length < 1 || length > 100000 {
		return 0
	}

	var k int = 2
	var profit [][]int = make([][]int, length)
	for i := range profit {
		profit[i] = make([]int, k+1)
	}
	var minArr []int = make([]int, k+1)

	for i := 1; i <= k; i++ {
		minArr[i] = prices[0]
	}

	for i := 1; i < length; i++ {
		for j := 1; j <= k; j++ {
			minArr[j] = min(prices[i]-profit[i][j-1], minArr[j])
			profit[i][j] = max(profit[i-1][j], prices[i]-minArr[j])
		}
	}
	return profit[length-1][k]
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
