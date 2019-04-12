package solution

func kthSmallest(matrix [][]int, k int) int {
	s := append([]int{}, matrix[0]...)
	for r := 1; r < len(matrix); r++ {
		// Find where this row should start being inserted.
		i := 0
		for ; i < len(s) && s[i] <= matrix[r][0]; i++ {
		}

		// If insertion point is at least k, we know the kth smallest.
		if i >= k {
			return s[k-1]
		}

		k -= i
		s = merge(s[i:], matrix[r])
	}
	return s[k-1]
}

func merge(a, b []int) []int {
	m := []int{}
	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			m = append(m, a[0])
			a = a[1:]
		} else {
			m = append(m, b[0])
			b = b[1:]
		}
	}
	if len(a) > 0 {
		m = append(m, a...)
	}
	if len(b) > 0 {
		m = append(m, b...)
	}
	return m
}
