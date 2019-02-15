package mxsq

import "math"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	w := len(matrix[0])

	// nz is the number of zeros above the current row.
	nz := make([]int, w)
	maxdim := 0
	for _, row := range matrix {
		for x := 0; x < w; x++ {
			// What's the largest square with lower left at x?
			minheight := math.MaxInt64
			dim := 0
			for i := x; i < w && row[i] == '1'; i++ {
				dim++
				minheight = min(minheight, nz[i]+1)
				if dim > minheight {
					break
				}
				maxdim = max(maxdim, dim)
			}
		}

		// Update zero counters.
		for x := 0; x < w; x++ {
			if row[x] == '1' {
				nz[x]++
			} else {
				nz[x] = 0
			}
		}
	}

	return maxdim * maxdim
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
