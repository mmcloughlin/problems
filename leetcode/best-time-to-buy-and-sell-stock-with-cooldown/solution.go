package solution

func maxProfit(prices []int) int {
	p2, p1, p := 0, 0, 0
	h := 0
	for t := 0; t+1 < len(prices); t++ {
		h = max(h, p2) + (prices[t+1] - prices[t])
		p2, p1, p = p1, p, max(h, p)
	}
	return p
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
