package uniqpaths

func uniquePaths(m int, n int) int {
	// we must make (m-1) steps right and (n-1) steps down.
	// so (m+n-2) steps total, the question boils down to "n choose k".
	if m < 2 {
		return 1
	}
	return choose(m-1, m+n-2)
}

func choose(k, n int) int {
	if k < n-k {
		k = n - k
	}
	c := 1
	for i := k + 1; i <= n; i++ {
		c *= i
	}
	for i := 2; i <= n-k; i++ {
		c /= i
	}
	return c
}
