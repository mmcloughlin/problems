package solution

import "sort"

// Unique returns the unique elements of x. May modify the input.
func Unique(x []int) []int {
	if len(x) == 0 {
		return x
	}

	// Sort.
	sort.Ints(x)

	// Unique.
	w := 1
	for r := 1; r < len(x); r++ {
		if x[r] != x[r-1] {
			x[w] = x[r]
			w++
		}
	}

	return x[:w]
}
