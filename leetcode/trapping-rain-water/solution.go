package solution

func trap(height []int) int {
	// running max from the left
	m := 0
	lmax := make([]int, len(height))
	for i, h := range height {
		lmax[i] = m
		m = max(m, h)
	}

	// running max from the right
	m = 0
	rmax := make([]int, len(height))
	for i := len(height) - 1; i >= 0; i-- {
		rmax[i] = m
		m = max(m, height[i])
	}

	// now count trapped water
	total := 0
	for i, h := range height {
		total += max(min(lmax[i], rmax[i])-h, 0)
	}

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	return -min(-a, -b)
}
