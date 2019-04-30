package solution

// merge merges b into a in place. Only the first m entries of a are to be
// considered, and the remaining positions will hold the result.
func merge(a, b []int, m int) {
	n := len(b)

	// Read and write indicies.
	ra, rb := m-1, n-1
	w := m + n - 1

	for ra >= 0 && rb >= 0 {
		if a[ra] >= b[rb] {
			a[w] = a[ra]
			ra--
		} else {
			a[w] = b[rb]
			rb--
		}
		w--
	}

	// Either ra >= 0 or rb >= 0
	// In the first case, there's nothing to be done.

	for rb >= 0 {
		a[w] = b[rb]
		rb--
		w--
	}
}
