package solution

func maxProfit(prices []int) int {
	profit := 0
	i := 0
	for i < len(prices) {
		j := i + 1
		for ; j < len(prices) && prices[j-1] <= prices[j]; j++ {
		}
		if j-i > 1 {
			profit += prices[j-1] - prices[i]
		}
		i = j
	}
	return profit
}
