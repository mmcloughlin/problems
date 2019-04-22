package solution

import "math"

const inf = math.MaxInt64

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	for r, c := m-2, n-2; r >= 0 || c >= 0; r, c = r-1, c-1 {
		// Set row.
		for c0 := n - 1; c0 > c && r >= 0 && c0 >= 0; c0-- {
			grid[r][c0] += min(at(grid, r+1, c0), at(grid, r, c0+1))
		}

		// Set col.
		for r0 := m - 1; r0 > r && r0 >= 0 && c >= 0; r0-- {
			grid[r0][c] += min(at(grid, r0+1, c), at(grid, r0, c+1))
		}

		// Set corner.
		if r >= 0 && c >= 0 {
			grid[r][c] += min(at(grid, r+1, c), at(grid, r, c+1))
		}
	}

	return grid[0][0]
}

func at(grid [][]int, r, c int) int {
	if 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0]) {
		return grid[r][c]
	}
	return inf
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
