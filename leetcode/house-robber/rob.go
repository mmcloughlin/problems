package rob

func rob(x []int) int {
	p, m := 0, 0
	for _, v := range x {
		p, m = m, max(v+p, m)
	}
	return m
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
