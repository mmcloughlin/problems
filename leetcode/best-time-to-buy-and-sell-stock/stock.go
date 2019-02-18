package stock

import "math"

func maxProfit(prices []int) int {
	maxprofit := 0
	maxrem := math.MinInt64
	for i := len(prices) - 1; i >= 0; i-- {
		p := prices[i]
		maxrem = max(p, maxrem)
		maxprofit = max(maxrem-p, maxprofit)
	}
	return maxprofit
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
