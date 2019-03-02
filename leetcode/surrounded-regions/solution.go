package solution

const (
	X = 'X'
	O = 'O'
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

func solve(board [][]byte) {
	// Dimensions.
	height := len(board)
	if height == 0 {
		return
	}
	width := len(board[0])

	// Maintain a matrix of points we've seen.
	seen := make([][]bool, height)
	for y := 0; y < height; y++ {
		seen[y] = make([]bool, width)
	}

	// Function to "explore" a region. This will collect all points in a region,
	// determine if it's surrounded, and if so set them all to X.
	explore := func(x, y int) {
		if seen[y][x] {
			return
		}
		seen[y][x] = true
		if board[y][x] == X {
			return
		}

		region := []point{}
		surrounded := true
		stack := []point{{x, y}}
		for len(stack) > 0 {
			top := len(stack) - 1
			p := stack[top]
			stack = stack[:top]

			region = append(region, p)

			for _, n := range p.neighbors() {
				if n.x < 0 || n.x >= width || n.y < 0 || n.y >= height {
					surrounded = false
					continue
				}
				if seen[n.y][n.x] || board[n.y][n.x] == X {
					continue
				}
				seen[n.y][n.x] = true
				stack = append(stack, n)
			}
		}

		if !surrounded {
			return
		}
		for _, p := range region {
			board[p.y][p.x] = X
		}
	}

	// Explore the entire board.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			explore(x, y)
		}
	}
}
