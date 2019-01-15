package maxarea

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxArea(height []int) int {
	// Indicies of progressive maximums from the left.
	max := -1
	lmax := []int{}
	for i, h := range height {
		if h > max {
			lmax = append(lmax, i)
			max = h
		}
	}

	// Indicies of progressive maximums from the right.
	max = -1
	rmax := []int{}
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			rmax = append(rmax, i)
			max = height[i]
		}
	}

	// Step through.
	maxa := 0
	for _, l := range lmax {
		for _, r := range rmax {
			if l >= r {
				continue
			}
			a := (r - l) * min(height[l], height[r])
			if a > maxa {
				maxa = a
			}
		}
	}

	return maxa
}
