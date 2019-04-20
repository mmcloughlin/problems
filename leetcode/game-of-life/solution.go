package solution

func gameOfLife(board [][]int) {
	// Determine dimensions.
	h := len(board)
	if h == 0 {
		return
	}
	w := len(board[0])

	// Count live neighbors.
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			for nr := r - 1; nr < r+2; nr++ {
				for nc := c - 1; nc < c+2; nc++ {
					if r == nr && c == nc {
						continue
					}
					if 0 <= nr && nr < h && 0 <= nc && nc < w {
						board[nr][nc] += (board[r][c] & 1) << 1
					}
				}
			}
		}
	}

	// Update states.
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			switch board[r][c] {
			case (2 << 1) | 1, (3 << 1) | 1, (3 << 1) | 0:
				board[r][c] = 1
			default:
				board[r][c] = 0
			}
		}
	}
}
