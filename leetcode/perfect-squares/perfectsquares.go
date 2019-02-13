package perfectsquares

import "math"

var cache = map[int]int{0: 0, 1: 1}

func numSquares(n int) int {
	// Look for a previously calculated result.
	if m, found := cache[n]; found {
		return m
	}

	m := math.MaxInt64

	// Squares less than or equal to n.
	for s := 1; s*s <= n; s++ {
		m = min(m, 1+numSquares(n-s*s))
	}

	cache[n] = m
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
