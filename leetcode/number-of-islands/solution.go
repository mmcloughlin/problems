package solution

const (
	land  = '1'
	water = '0'
)

type point struct {
	x, y int
}

func (p point) neighbors() []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

func numIslands(grid [][]byte) int {
	// Determine grid dimensions.
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])

	// Boolean array for positions we've seen.
	seen := make([][]bool, height)
	for y := 0; y < height; y++ {
		seen[y] = make([]bool, width)
	}

	// Explore function to mark all positions of an island seen.
	explore := func(origin point) {
		stack := []point{origin}
		for len(stack) > 0 {
			top := len(stack) - 1
			p := stack[top]
			stack = stack[:top]

			seen[p.y][p.x] = true

			for _, n := range p.neighbors() {
				if n.x < 0 || n.x >= width || n.y < 0 || n.y >= height {
					continue
				}
				if !seen[n.y][n.x] && grid[n.y][n.x] == land {
					stack = append(stack, n)
				}
			}
		}
	}

	// Count islands.
	n := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if !seen[y][x] && grid[y][x] == land {
				n++
				explore(point{x, y})
			}
		}
	}

	return n
}
