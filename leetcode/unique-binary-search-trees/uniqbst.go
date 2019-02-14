package uniqbst

// cache stores memoized numTrees results.
var cache = map[int]int{0: 1}

// numTrees returns the number of BSTs containing values 1...n.
func numTrees(n int) int {
	if c, found := cache[n]; found {
		return c
	}

	c := 0
	for i := 1; i <= n; i++ {
		c += numTrees(i-1) * numTrees(n-i)
	}

	cache[n] = c
	return c
}
