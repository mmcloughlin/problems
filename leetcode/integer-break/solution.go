package solution

func integerBreak(n int) int {
	m := 0
	for a := 1; a < n; a++ {
		p := maxSumPartition(a) * maxSumPartition(n-a)
		if p > m {
			m = p
		}
	}
	return m
}

var cache = map[int]int{}

func maxSumPartition(n int) int {
	if p, ok := cache[n]; ok {
		return p
	}

	m := n
	for a := 1; a < n; a++ {
		b := n - a
		p := maxSumPartition(a) * maxSumPartition(b)
		if p > m {
			m = p
		}
	}

	cache[n] = m
	return m
}
